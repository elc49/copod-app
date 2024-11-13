-- name: CreateSupportDoc :one
INSERT INTO support_docs (
  govt_id, email, wallet_address
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetEmailSupportDoc :one
SELECT * FROM support_docs
WHERE email = $1 LIMIT 1;

-- name: UpdateEmailSupportDoc :one
UPDATE support_docs SET govt_id = $1, verification = $2
WHERE email = $3
RETURNING *;
