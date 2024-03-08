package db

import (
	"context"
	"testing"
	"time"

	"github.com/papaya147/money-tracker/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomExpenditure(t *testing.T) Expenditure {
	category := createRandomExpenditureCategory(t)

	arg := CreateExpenditureParams{
		Paisa:      util.RandomInt32(0, 10000),
		Categoryid: category.Name,
	}

	expenditure, err := testQueries.CreateExpenditure(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, expenditure)
	require.Equal(t, arg.Paisa, expenditure.Paisa)
	require.Equal(t, arg.Categoryid, expenditure.Categoryid)
	require.NotZero(t, expenditure.ID)
	require.NotZero(t, expenditure.Createdat)
	require.NotZero(t, expenditure.Updatedat)

	return expenditure
}

func TestCreateExpenditure(t *testing.T) {
	createRandomExpenditure(t)
}

func TestGetExpenditures(t *testing.T) {
	expenditure := createRandomExpenditure(t)

	arg := int32(0)

	_, err := testQueries.GetExpenditures(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, expenditure)
}

func TestUpdateExpenditure(t *testing.T) {
	expenditure1 := createRandomExpenditure(t)
	category1 := createRandomExpenditureCategory(t)

	arg := UpdateExpenditureParams{
		Paisa:      util.RandomInt32(0, 10000),
		Categoryid: category1.Name,
		ID:         expenditure1.ID,
	}

	expenditure2, err := testQueries.UpdateExpenditure(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, expenditure2)
	require.Equal(t, arg.Categoryid, expenditure2.Categoryid)
	require.Equal(t, arg.Paisa, expenditure2.Paisa)
	require.Equal(t, arg.Categoryid, expenditure2.Categoryid)
	require.WithinDuration(t, expenditure1.Createdat, expenditure2.Createdat, time.Second)
}

func TestDeleteExpenditure(t *testing.T) {

}
