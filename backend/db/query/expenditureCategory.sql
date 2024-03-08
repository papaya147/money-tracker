-- name: CreateExpenditureCategory :one
INSERT INTO expenditureCategory (name)
VALUES ($1)
RETURNING *;
-- name: GetExpenditureCategories :many
SELECT *
FROM expenditureCategory;
-- name: UpdateExpenditureCategory :one
UPDATE expenditureCategory
SET name = $1,
    updatedAt = NOW()
WHERE name = $2
RETURNING *;
-- name: DeleteExpenditureCategory :one
DELETE FROM expenditureCategory
WHERE name = $1
RETURNING *;