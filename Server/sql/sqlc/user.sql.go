// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package sql

import (
	"context"

	"github.com/google/uuid"
)

const countUsers = `-- name: CountUsers :one
SELECT COUNT(*) FROM users
`

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUsers)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email, firstname, lastname
) VALUES (
  $1, $2, $3
) RETURNING id, firstname, lastname, email, email_onboarded, created_at, updated_at
`

type CreateUserParams struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Firstname, arg.Lastname)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.EmailOnboarded,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, firstname, lastname, email, email_onboarded, created_at, updated_at FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.EmailOnboarded,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserEmailOnboardByID = `-- name: UpdateUserEmailOnboardByID :one
UPDATE users SET email_onboarded = $1
WHERE id = $2
RETURNING id, firstname, lastname, email, email_onboarded, created_at, updated_at
`

type UpdateUserEmailOnboardByIDParams struct {
	EmailOnboarded bool      `json:"email_onboarded"`
	ID             uuid.UUID `json:"id"`
}

func (q *Queries) UpdateUserEmailOnboardByID(ctx context.Context, arg UpdateUserEmailOnboardByIDParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserEmailOnboardByID, arg.EmailOnboarded, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.EmailOnboarded,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
