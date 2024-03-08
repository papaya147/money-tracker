package category

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/papaya147/money-tracker/backend/util"
)

// update godoc
// @Summary      Delete an expenditure category's
// @Description  Delete an expenditure category's, this API will return an error if the category does not exist.
// @Tags         expenditure category
// @Param        category-name path string true "Category Name" maxlength(32)
// @Success      200  {object} categoryOutput
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /expenditure/category/{category-name} [delete]
func (c *Controller) delete(w http.ResponseWriter, r *http.Request) {
	requestPayload := deleteCategoryInput{
		Name: chi.URLParam(r, "category-name"),
	}

	if err := util.ValidateRequest(requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	queryArg := requestPayload.Name

	category, err := c.store.DeleteExpenditureCategory(r.Context(), queryArg)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.ErrorJson(w, util.ErrEmptyResult)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, categoryOutput{
		Name:      category.Name,
		CreatedAt: category.Createdat.UnixMilli(),
		UpdatedAt: category.Updatedat.UnixMilli(),
	})
}
