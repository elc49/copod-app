-- name: CreateTitle :one
INSERT INTO titles (
  title, email
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetTitleByEmail :one
SELECT * FROM titles
WHERE email = $1 LIMIT 1;

-- name: UpdateTitleByEmail :one
UPDATE titles SET title = $1, verification = $2
WHERE email = $3
RETURNING *;

-- name: UpdateTitleVerificationById :one
UPDATE titles SET verification = $1
WHERE id = $2
RETURNING *;

-- name: GetTitlesByEmail :many
SELECT * FROM titles
WHERE email = $1;
