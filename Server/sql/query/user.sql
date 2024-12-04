-- name: CreateUser :one
INSERT INTO users (
  email, firstname, lastname, govt_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateUserSupportDocByEmail :one
UPDATE support_docs SET govt_id = $1, verification = $2
WHERE email = $3
RETURNING *;
