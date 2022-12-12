-- name: CreateCompany :one
INSERT INTO companies (
  owner,
  headquarters,
  founded 
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetCompany :one
SELECT * FROM companies
WHERE id = $1 LIMIT 1;

-- name: ListCompanies :many
SELECT * FROM companies
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCompany :one
UPDATE companies
SET headquarters = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCompany :exec
DELETE FROM companies
WHERE id = $1;