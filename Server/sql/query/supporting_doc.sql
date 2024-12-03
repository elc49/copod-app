-- name: CreateSupportDoc :one
INSERT INTO support_docs (
  email, govt_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetSupportDocByEmail :one
SELECT * FROM support_docs
WHERE email = $1 LIMIT 1;

-- name: UpdateSupportDocByEmail :one
UPDATE support_docs SET govt_id = $1, verification = $2
WHERE email = $3
RETURNING *;

-- name: GetSupportingDocsByVerification :many
SELECT * FROM support_docs
WHERE verification = $1;
