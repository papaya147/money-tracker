package expenditure

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/papaya147/money-tracker/backend/util"
)

// get godoc
// @Summary      Get all expenditures, 20 at a time
// @Description  Get all expenditures, this API will return an error if the offset is empty.
// @Tags         expenditure
// @Param        page query int true "Page" int
// @Success      200  {object} []expenditureOutput
// @Failure      400  {object} util.ErrorModel
// @Failure      500  {object} util.ErrorModel
// @Router       /expenditure [get]
func (c *Controller) get(w http.ResponseWriter, r *http.Request) {
	requestPayload := getExpendituresInput{
		Page: r.URL.Query().Get("page"),
	}

	queryArg, _ := strconv.Atoi(requestPayload.Page)

	exps, err := c.store.GetExpenditures(r.Context(), 20*int32(queryArg))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.ErrorJson(w, util.ErrEmptyResult)
			return
		}
		util.ErrorJson(w, util.ErrDatabase)
		return
	}

	out := make([]expenditureOutput, len(exps))
	for i, exp := range exps {
		out[i] = expenditureOutput{
			Id:        exp.ID,
			Paisa:     exp.Paisa,
			Category:  exp.Categoryid,
			CreatedAt: exp.Createdat.UnixMilli(),
			UpdatedAt: exp.Updatedat.UnixMilli(),
		}
	}

	util.WriteJson(w, http.StatusOK, out)
}
