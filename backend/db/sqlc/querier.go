// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateExpenditure(ctx context.Context, arg CreateExpenditureParams) (Expenditure, error)
	CreateExpenditureCategory(ctx context.Context, name string) (Expenditurecategory, error)
	DeleteExpenditure(ctx context.Context, id uuid.UUID) (Expenditure, error)
	DeleteExpenditureCategory(ctx context.Context, name string) (Expenditurecategory, error)
	GetExpenditureById(ctx context.Context, id uuid.UUID) (Expenditure, error)
	GetExpenditureCategories(ctx context.Context) ([]Expenditurecategory, error)
	GetExpenditureCategoryById(ctx context.Context, id uuid.UUID) (Expenditurecategory, error)
	GetExpenditures(ctx context.Context, offset int32) ([]GetExpendituresRow, error)
	UpdateExpenditure(ctx context.Context, arg UpdateExpenditureParams) (Expenditure, error)
	UpdateExpenditureCategory(ctx context.Context, arg UpdateExpenditureCategoryParams) (Expenditurecategory, error)
}

var _ Querier = (*Queries)(nil)
