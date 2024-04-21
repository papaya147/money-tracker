-- name: CreateExpenditure :one
INSERT INTO expenditure (paisa, categoryId, createdAt)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetExpenditures :many
SELECT e.*,
    c.name AS categoryName
FROM expenditure e
    INNER JOIN expenditureCategory c ON e.categoryId = c.id
LIMIT 20 OFFSET $1;
-- name: UpdateExpenditure :one
UPDATE expenditure
SET paisa = $1,
    categoryId = $2,
    createdAt = $3,
    updatedAt = NOW()
WHERE id = $4
RETURNING *;
-- name: DeleteExpenditure :one
DELETE FROM expenditure
WHERE id = $1
RETURNING *;