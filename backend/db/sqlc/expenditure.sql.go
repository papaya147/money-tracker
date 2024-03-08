// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: expenditure.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createExpenditure = `-- name: CreateExpenditure :one
INSERT INTO expenditure (paisa, categoryId)
VALUES ($1, $2)
RETURNING id, paisa, categoryid, createdat, updatedat
`

type CreateExpenditureParams struct {
	Paisa      int32  `json:"paisa"`
	Categoryid string `json:"categoryid"`
}

func (q *Queries) CreateExpenditure(ctx context.Context, arg CreateExpenditureParams) (Expenditure, error) {
	row := q.db.QueryRow(ctx, createExpenditure, arg.Paisa, arg.Categoryid)
	var i Expenditure
	err := row.Scan(
		&i.ID,
		&i.Paisa,
		&i.Categoryid,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const deleteExpenditure = `-- name: DeleteExpenditure :one
DELETE FROM expenditure
WHERE id = $1
RETURNING id, paisa, categoryid, createdat, updatedat
`

func (q *Queries) DeleteExpenditure(ctx context.Context, id uuid.UUID) (Expenditure, error) {
	row := q.db.QueryRow(ctx, deleteExpenditure, id)
	var i Expenditure
	err := row.Scan(
		&i.ID,
		&i.Paisa,
		&i.Categoryid,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const getExpenditures = `-- name: GetExpenditures :many
SELECT id, paisa, categoryid, createdat, updatedat
FROM expenditure
LIMIT 20 OFFSET $1
`

func (q *Queries) GetExpenditures(ctx context.Context, offset int32) ([]Expenditure, error) {
	rows, err := q.db.Query(ctx, getExpenditures, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Expenditure{}
	for rows.Next() {
		var i Expenditure
		if err := rows.Scan(
			&i.ID,
			&i.Paisa,
			&i.Categoryid,
			&i.Createdat,
			&i.Updatedat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExpenditure = `-- name: UpdateExpenditure :one
UPDATE expenditure
SET paisa = $1,
    categoryId = $2,
    updatedAt = NOW()
WHERE id = $3
RETURNING id, paisa, categoryid, createdat, updatedat
`

type UpdateExpenditureParams struct {
	Paisa      int32     `json:"paisa"`
	Categoryid string    `json:"categoryid"`
	ID         uuid.UUID `json:"id"`
}

func (q *Queries) UpdateExpenditure(ctx context.Context, arg UpdateExpenditureParams) (Expenditure, error) {
	row := q.db.QueryRow(ctx, updateExpenditure, arg.Paisa, arg.Categoryid, arg.ID)
	var i Expenditure
	err := row.Scan(
		&i.ID,
		&i.Paisa,
		&i.Categoryid,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}