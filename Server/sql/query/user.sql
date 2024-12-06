-- name: CreateUser :one
INSERT INTO users (
  email, firstname, lastname
) VALUES (
  $1, $2, $3
) RETURNING *;
