-- name: CreateEarlySignup :one
INSERT INTO early_signups (
  email
) VALUES (
  $1
) RETURNING *;

-- name: GetEarlySignupByEmail :one
SELECT email FROM early_signups
WHERE email = $1 LIMIT 1;
