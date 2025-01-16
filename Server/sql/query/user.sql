-- name: CreateUser :one
INSERT INTO users (
  email, firstname, lastname
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: UpdateUserEmailOnboardByID :one
UPDATE users SET email_onboarded = $1
WHERE id = $2
RETURNING *;
