-- name: CreateUpload :one
INSERT INTO uploads (
  type, uri, verification, wallet_address
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUpload :one
SELECT * FROM uploads
WHERE type = $1 AND wallet_address = $2;

-- name: ClearTestUploads :exec
DELETE FROM uploads;
