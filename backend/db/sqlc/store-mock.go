package db

import (
	"context"

	"github.com/google/uuid"
)

type mockStore struct {
}

// GetExpenditureById implements Store.
func (m *mockStore) GetExpenditureById(ctx context.Context, id uuid.UUID) (Expenditure, error) {
	return Expenditure{}, nil
}

// GetExpenditureCategoryById implements Store.
func (m *mockStore) GetExpenditureCategoryById(ctx context.Context, id uuid.UUID) (Expenditurecategory, error) {
	return Expenditurecategory{}, nil
}

// GetExpenditures implements Store.
func (m *mockStore) GetExpenditures(ctx context.Context, offset int32) ([]GetExpendituresRow, error) {
	return nil, nil
}

// CreateExpenditure implements Store.
func (m *mockStore) CreateExpenditure(ctx context.Context, arg CreateExpenditureParams) (Expenditure, error) {
	return Expenditure{}, nil
}

// DeleteExpenditure implements Store.
func (m *mockStore) DeleteExpenditure(ctx context.Context, id uuid.UUID) (Expenditure, error) {
	return Expenditure{}, nil
}

// UpdateExpenditure implements Store.
func (m *mockStore) UpdateExpenditure(ctx context.Context, arg UpdateExpenditureParams) (Expenditure, error) {
	return Expenditure{}, nil
}

// CreateExpenditureCategory implements Store.
func (m *mockStore) CreateExpenditureCategory(ctx context.Context, name string) (Expenditurecategory, error) {
	return Expenditurecategory{}, nil
}

// DeleteExpenditureCategory implements Store.
func (m *mockStore) DeleteExpenditureCategory(ctx context.Context, name string) (Expenditurecategory, error) {
	return Expenditurecategory{}, nil
}

// GetExpenditureCategories implements Store.
func (m *mockStore) GetExpenditureCategories(ctx context.Context) ([]Expenditurecategory, error) {
	return nil, nil
}

// UpdateExpenditureCategory implements Store.
func (m *mockStore) UpdateExpenditureCategory(ctx context.Context, arg UpdateExpenditureCategoryParams) (Expenditurecategory, error) {
	return Expenditurecategory{}, nil
}

func NewMockStore() Store {
	return &mockStore{}
}
