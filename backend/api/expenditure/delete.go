package expenditure

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/papaya147/money-tracker/backend/util"
)

// delete godoc
// @Summary      Delete an expenditure
// @Description  Delete an expenditure, this API will return an error if the expenditure is not found.
// @Tags         expenditure
// @Param        expenditure-id path string true "Expenditure Id" uuid
// @Success      200  {object} expenditureOutput
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /expenditure/{expenditure-id} [delete]
func (c *Controller) delete(w http.ResponseWriter, r *http.Request) {
	requsetPayload := deleteExpenditureInput{
		Id: chi.URLParam(r, "expenditure-id"),
	}

	if err := util.ValidateRequest(requsetPayload); err != nil {
		util.ErrorJson(w, err)
		return
	}

	queryArg := uuid.MustParse(requsetPayload.Id)
	exp, err := c.store.DeleteExpenditure(r.Context(), queryArg)
	if err != nil {
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
