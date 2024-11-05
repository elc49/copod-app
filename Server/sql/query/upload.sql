-- name: CreateUpload :one
INSERT INTO uploads (
  type, uri, verification, wallet_address
) VALUES (
  $1, $2, $3, $4
) RETURNING *;
