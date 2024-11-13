-- name: CreateUser :one
INSERT INTO users (
  email, wallet_address, govt_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE wallet_address = $1;

-- name: UpdateUser :one
UPDATE users SET firstname = $1, lastname = $2
WHERE email = $3
RETURNING  *;
