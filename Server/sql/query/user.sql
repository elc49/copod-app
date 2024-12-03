-- name: CreateUser :one
INSERT INTO users (
  email, firstname, lastname, govt_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;
