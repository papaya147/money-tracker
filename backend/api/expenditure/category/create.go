package category

import (
	"net/http"

	db "github.com/papaya147/money-tracker/backend/db/sqlc"
	"github.com/papaya147/money-tracker/backend/util"
)

// create godoc
// @Summary      Create a new expenditure category
// @Description  Create a new expenditure category, this API will return an error if the category has already been created.
// @Tags         expenditure category
// @Param        request-body body createCategoryInput true "json"
// @Success      200  {object} categoryOutput
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /expenditure/category [post]
func (c *Controller) create(w http.ResponseWriter, r *http.Request) {
	var requestPayload createCategoryInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	queryArg := requestPayload.Name
	category, err := c.store.CreateExpenditureCategory(r.Context(), queryArg)
	if err != nil {
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
