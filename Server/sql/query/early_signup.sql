-- name: CreateEarlySignup :one
INSERT INTO early_signups (
  email
) VALUES (
  $1
) RETURNING *;

-- name: GetEarlySignupByEmail :one
SELECT * FROM early_signups
WHERE email = $1;

-- name: OnboardEarlySignup :one
UPDATE early_signups SET onboarded = $1
WHERE email = $2
RETURNING *;
