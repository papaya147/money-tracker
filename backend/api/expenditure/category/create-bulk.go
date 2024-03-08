package category

import (
	"net/http"

	db "github.com/papaya147/money-tracker/backend/db/sqlc"
	"github.com/papaya147/money-tracker/backend/util"
)

// createBulk godoc
// @Summary      Create a bulk of new expenditure categories
// @Description  Create a bulk of new expenditure categories, this API will return an error if any category already exists.
// @Tags         expenditure category
// @Param        request-body body createBulkCategoryInput true "json"
// @Success      200  {object} []categoryOutput
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /expenditure/category/bulk [post]
func (c *Controller) createBulk(w http.ResponseWriter, r *http.Request) {
	var requestPayload createBulkCategoryInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	out := make([]categoryOutput, len(requestPayload.Categories))
	for i, category := range requestPayload.Categories {
		category, err := c.store.CreateExpenditureCategory(r.Context(), category)
		if err != nil {
			if db.ErrorCode(err) == db.UniqueViolation {
				util.ErrorJson(w, util.ErrDuplicateEntry)
				return
			}
			util.ErrorJson(w, util.ErrDatabase)
			return
		}

		out[i] = categoryOutput{
			Name:      category.Name,
			CreatedAt: category.Createdat.UnixMilli(),
			UpdatedAt: category.Updatedat.UnixMilli(),
		}
	}

	util.WriteJson(w, http.StatusOK, out)
}
