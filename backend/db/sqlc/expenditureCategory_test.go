package db

import (
	"context"
	"testing"
	"time"

	"github.com/papaya147/money-tracker/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomExpenditureCategory(t *testing.T) Expenditurecategory {
	arg := util.RandomString(10)

	category, err := testQueries.CreateExpenditureCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, arg, category.Name)
	require.NotZero(t, category.ID)
	require.NotZero(t, category.Createdat)
	require.NotZero(t, category.Updatedat)

	return category
}

func TestCreateExpenditureCategory(t *testing.T) {
	createRandomExpenditureCategory(t)
}

func TestGetExpenditureCategories(t *testing.T) {
	category := createRandomExpenditureCategory(t)

	categories, err := testQueries.GetExpenditureCategories(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, categories)
	require.Contains(t, categories, category)
}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomExpenditureCategory(t)

	arg := UpdateExpenditureCategoryParams{
		Name:   util.RandomString(10),
		Name_2: category1.Name,
	}

	category2, err := testQueries.UpdateExpenditureCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.ID, category2.ID)
	require.NotEqual(t, category1.Name, category2.Name)
	require.WithinDuration(t, category1.Createdat, category2.Createdat, time.Second)
}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomExpenditureCategory(t)

	arg := category1.Name

	category2, err := testQueries.DeleteExpenditureCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.WithinDuration(t, category1.Createdat, category2.Createdat, time.Second)
	require.WithinDuration(t, category1.Updatedat, category2.Updatedat, time.Second)

	categories, err := testQueries.GetExpenditureCategories(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, categories)
	require.NotContains(t, categories, category1)
}

func TestGetExpenditureCategoryById(t *testing.T) {
	category1 := createRandomExpenditureCategory(t)

	arg := category1.ID

	category2, err := testQueries.GetExpenditureCategoryById(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.WithinDuration(t, category1.Createdat, category2.Createdat, time.Second)
	require.WithinDuration(t, category1.Updatedat, category2.Updatedat, time.Second)
}
