-- name: CreateUpload :one
INSERT INTO uploads (
  type, title_doc, govt_id, email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: ClearTestUploads :exec
DELETE FROM uploads;
