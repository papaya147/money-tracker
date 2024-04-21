package category

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/money-tracker/backend/db/sqlc"
	"github.com/papaya147/money-tracker/backend/util"
)

// update godoc
// @Summary      Update an expenditure category's name
// @Description  Update an expenditure category's name, this API will return an error if the category does not exist.
// @Tags         expenditure category
// @Param        category-name path string true "Category Name" maxlength(32)
// @Param        request-body body updateCategoryInput true "json"
// @Success      200  {object} categoryOutput
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /expenditure/category/{category-name} [put]
func (c *Controller) update(w http.ResponseWriter, r *http.Request) {
	var requestPayload updateCategoryInput
	if err := util.ReadJson(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}
	requestPayload.OldName = chi.URLParam(r, "category-name")

	if err := util.ValidateRequest(requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	queryArg := db.UpdateExpenditureCategoryParams{
		Name:   requestPayload.NewName,
		Name_2: requestPayload.OldName,
	}
	category, err := c.store.UpdateExpenditureCategory(r.Context(), queryArg)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.ErrorJson(w, util.ErrEmptyResult)
			return
		}
		if db.ErrorCode(err) == db.UniqueViolation {
			util.ErrorJson(w, util.ErrDuplicateEntry)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, categoryOutput{
		Id:        category.ID,
		Name:      category.Name,
		CreatedAt: category.Createdat.UnixMilli(),
		UpdatedAt: category.Updatedat.UnixMilli(),
	})
}
