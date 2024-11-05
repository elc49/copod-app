-- name: CreateUpload :execresult
INSERT INTO uploads (
  email, public_address, type, uri, verification
) VALUES (
  $1, $2, $3, $4, $5
);
