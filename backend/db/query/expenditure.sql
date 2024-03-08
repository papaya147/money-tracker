-- name: CreateExpenditure :one
INSERT INTO expenditure (paisa, categoryId)
VALUES ($1, $2)
RETURNING *;
-- name: GetExpenditures :many
SELECT *
FROM expenditure
LIMIT 20 OFFSET $1;
-- name: UpdateExpenditure :one
UPDATE expenditure
SET paisa = $1,
    categoryId = $2,
    updatedAt = NOW()
WHERE id = $3
RETURNING *;
-- name: DeleteExpenditure :one
DELETE FROM expenditure
WHERE id = $1
RETURNING *;