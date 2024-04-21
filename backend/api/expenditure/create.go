package expenditure

import (
	"net/http"
	"time"

	db "github.com/papaya147/money-tracker/backend/db/sqlc"
	"github.com/papaya147/money-tracker/backend/util"
)

// create godoc
// @Summary      Create a new expenditure
// @Description  Create a new expenditure, this API will return an error if the category is not found.
// @Tags         expenditure
// @Param        request-body body createExpenditureInput true "json"
// @Success      200  {object} expenditureOutput
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /expenditure [post]
func (c *Controller) create(w http.ResponseWriter, r *http.Request) {
	var requestPayload createExpenditureInput
	if err := util.ReadJsonAndValidate(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	if requestPayload.UnixMillis == 0 {
		requestPayload.UnixMillis = time.Now().UnixMilli()
	}

	queryArg := db.CreateExpenditureParams{
		Paisa:      requestPayload.Paisa,
		Categoryid: requestPayload.CategoryId,
		Createdat:  time.UnixMilli(requestPayload.UnixMillis),
	}

	exp, err := c.store.CreateExpenditure(r.Context(), queryArg)
	if err != nil {
		if db.ErrorCode(err) == db.ForeignKeyViolation {
			util.ErrorJson(w, util.ErrCategoryNotFound)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	query2Arg := exp.Categoryid
	category, err := c.store.GetExpenditureCategoryById(r.Context(), query2Arg)
	if err != nil {
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, expenditureOutput{
		Id:        exp.ID,
		Paisa:     exp.Paisa,
		Category:  category.Name,
		CreatedAt: exp.Createdat.UnixMilli(),
		UpdatedAt: exp.Updatedat.UnixMilli(),
	})
}
