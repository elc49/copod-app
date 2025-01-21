-- name: CreateOnboarding :one
INSERT INTO onboardings (
  title_id, support_doc_id, display_picture_id, email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetOnboardingByEmailAndVerification :one
SELECT * FROM onboardings
WHERE email = $1 AND verification = $2;

-- name: GetOnboardingByID :one
SELECT * FROM onboardings
WHERE id = $1;

-- name: UpdateOnboardingVerificationByID :one
UPDATE onboardings SET verification = $1
WHERE id = $2
RETURNING *;

-- name: GetOnboardingsByStatus :many
SELECT * FROM onboardings
WHERE verification = $1;
