-- name: CreateTitle :one
INSERT INTO title_deeds (
  url, email, support_doc_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTitleByEmail :one
SELECT * FROM title_deeds
WHERE email = $1 LIMIT 1;

-- name: UpdateTitleByID :one
UPDATE title_deeds SET url = $1, verification = $2
WHERE id = $3
RETURNING *;

-- name: GetTitlesByEmail :many
SELECT * FROM title_deeds
WHERE email = $1;
