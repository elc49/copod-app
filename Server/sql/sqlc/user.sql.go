// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package sql

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email, firstname, lastname, govt_id
) VALUES (
  $1, $2, $3, $4
) RETURNING id, firstname, lastname, govt_id, email, created_at, updated_at
`

type CreateUserParams struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	GovtID    string `json:"govt_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		arg.GovtID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.GovtID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserSupportDocByEmail = `-- name: UpdateUserSupportDocByEmail :one
UPDATE support_docs SET govt_id = $1, verification = $2
WHERE email = $3
RETURNING id, govt_id, verification, email, created_at, updated_at
`

type UpdateUserSupportDocByEmailParams struct {
	GovtID       string `json:"govt_id"`
	Verification string `json:"verification"`
	Email        string `json:"email"`
}

func (q *Queries) UpdateUserSupportDocByEmail(ctx context.Context, arg UpdateUserSupportDocByEmailParams) (SupportDoc, error) {
	row := q.db.QueryRowContext(ctx, updateUserSupportDocByEmail, arg.GovtID, arg.Verification, arg.Email)
	var i SupportDoc
	err := row.Scan(
		&i.ID,
		&i.GovtID,
		&i.Verification,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
