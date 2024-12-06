-- name: CreateSupportDoc :one
INSERT INTO support_docs (
  email, url
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetSupportDocByEmail :one
SELECT * FROM support_docs
WHERE email = $1 LIMIT 1;

-- name: GetSupportingDocsByVerification :many
SELECT * FROM support_docs
WHERE verification = $1;

-- name: GetSupportingDocById :one
SELECT * FROM support_docs
WHERE id = $1;

-- name: UpdateSupportDocVerificationById :one
UPDATE support_docs SET verification = $1
WHERE id = $2
RETURNING *;
