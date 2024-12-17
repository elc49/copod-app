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

-- name: GetSupportDocByID :one
SELECT * FROM support_docs
WHERE id = $1;

-- name: UpdateSupportDocByID :one
UPDATE support_docs SET url = $1, verification = $2
WHERE id = $3
RETURNING *;

-- name: UpdateSupportDocVerificationByID :one
UPDATE support_docs SET verification = $1
WHERE id = $2
RETURNING *;
