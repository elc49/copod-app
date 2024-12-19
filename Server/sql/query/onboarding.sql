-- name: CreateOnboarding :one
INSERT INTO onboardings (
  title_id, support_doc_id, display_picture_id, email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetOnboardingByEmail :one
SELECT * FROM onboardings
WHERE email = $1;

-- name: GetOnboardingByID :one
SELECT * FROM onboardings
WHERE id = $1;
