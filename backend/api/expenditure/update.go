package expenditure

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	db "github.com/papaya147/money-tracker/backend/db/sqlc"
	"github.com/papaya147/money-tracker/backend/util"
)

// update godoc
// @Summary      Update an expenditure
// @Description  Update an expenditure, this API will return an error if the category is not found.
// @Tags         expenditure
// @Param        expenditure-id path string true "Expenditure Id" uuid
// @Param        request-body body updateExpenditureInput true "json"
// @Success      200  {object} expenditureOutput
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /expenditure/{expenditure-id} [put]
func (c *Controller) update(w http.ResponseWriter, r *http.Request) {
	var requestPayload updateExpenditureInput
	if err := util.ReadJson(w, r, &requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	requestPayload.Id = chi.URLParam(r, "expenditure-id")

	if err := util.ValidateRequest(requestPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	queryArg := db.UpdateExpenditureParams{
		Paisa:      requestPayload.Paisa,
		Categoryid: requestPayload.Category,
		ID:         uuid.MustParse(requestPayload.Id),
	}
	exp, err := c.store.UpdateExpenditure(r.Context(), queryArg)
	if err != nil {
		if db.ErrorCode(err) == db.ForeignKeyViolation {
			util.ErrorJson(w, util.ErrCategoryNotFound)
			return
		}
		if errors.Is(err, pgx.ErrNoRows) {
			util.ErrorJson(w, util.ErrEmptyResult)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	util.WriteJson(w, http.StatusOK, expenditureOutput{
		Id:        exp.ID,
		Paisa:     exp.Paisa,
		Category:  exp.Categoryid,
		CreatedAt: exp.Createdat.UnixMilli(),
		UpdatedAt: exp.Updatedat.UnixMilli(),
	})
}
