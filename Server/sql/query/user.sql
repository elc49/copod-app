-- name: CreateUser :one
INSERT INTO users (
  email, wallet_address
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE wallet_address = $1;

-- name: UpdateUser :one
UPDATE users SET firstname = $1, lastname = $2, govt_id = $3
WHERE email = $4
RETURNING  *;
