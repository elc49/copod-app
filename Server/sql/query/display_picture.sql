-- name: CreateDisplayPicture :one
INSERT INTO display_pictures (
  url, email, support_doc_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetDisplayPictureByID :one
SELECT * FROM display_pictures
WHERE id = $1;

-- name: UpdateDisplayPictureByID :one
UPDATE display_pictures SET url = $1, verification = $2
WHERE id = $3
RETURNING *;

-- name: UpdateDisplayPictureVerificationByID :one
UPDATE display_pictures SET verification = $1
WHERE id = $2
RETURNING *;
