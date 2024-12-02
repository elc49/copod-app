-- name: CreateTitle :one
INSERT INTO titles (
  title, email, wallet_address
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetEmailTitle :one
SELECT * FROM titles
WHERE email = $1 LIMIT 1;

-- name: UpdateEmailTitle :one
UPDATE titles SET title = $1, verification = $2
WHERE email = $3
RETURNING *;

-- name: UpdateTitleVerification :one
UPDATE titles SET verification = $1
WHERE id = $2
RETURNING *;
