-- name: CreateUser :one
INSERT INTO users (
  email
) VALUES (
  $1
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUserByEmail :one
UPDATE users SET firstname = $1, lastname = $2, govt_id = $3
WHERE email = $4
RETURNING  *;
