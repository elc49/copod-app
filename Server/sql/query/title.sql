-- name: CreateTitle :one
INSERT INTO titles (
  title, email, wallet_address
) VALUES (
  $1, $2, $3
) RETURNING *;
