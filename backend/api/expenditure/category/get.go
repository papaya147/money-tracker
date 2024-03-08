package category

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/papaya147/money-tracker/backend/util"
)

// create godoc
// @Summary      Get all expenditure categories
// @Description  Get all expenditure categories, this API will return an error if none have been created.
// @Tags         expenditure category
// @Success      200  {object} []string
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /expenditure/category [get]
func (c *Controller) get(w http.ResponseWriter, r *http.Request) {
	categories, err := c.store.GetExpenditureCategories(r.Context())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.ErrorJson(w, util.ErrEmptyResult)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	out := make([]string, len(categories))
	for i, category := range categories {
		out[i] = category.Name
	}

	util.WriteJson(w, http.StatusOK, out)
}
