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

-- name: GetTitlesByEmailAndVerification :many
SELECT * FROM title_deeds
WHERE email = $1 AND verification = $2;

-- name: GetTitleByID :one
SELECT * FROM title_deeds
WHERE id = $1;

-- name: UpdateTitleVerificationByID :one
UPDATE title_deeds SET verification = $1
WHERE id = $2
RETURNING *;
