-- name: CreateTitle :one
INSERT INTO title_deeds (
  url, title, email, support_doc_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetTitleByEmail :one
SELECT * FROM title_deeds
WHERE email = $1 LIMIT 1;

-- name: UpdateTitleVerificationById :one
UPDATE title_deeds SET verification = $1
WHERE id = $2
RETURNING *;

-- name: GetTitlesByEmail :many
SELECT * FROM title_deeds
WHERE email = $1;
