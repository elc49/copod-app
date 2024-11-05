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
  email, wallet_address
) VALUES (
  $1, $2
) RETURNING id, firstname, lastname, govt_id, email, wallet_address, created_at, updated_at
`

type CreateUserParams struct {
	Email         string `json:"email"`
	WalletAddress string `json:"wallet_address"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.WalletAddress)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.GovtID,
		&i.Email,
		&i.WalletAddress,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, firstname, lastname, govt_id, email, wallet_address, created_at, updated_at FROM users
WHERE wallet_address = $1
`

func (q *Queries) GetUser(ctx context.Context, walletAddress string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, walletAddress)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.GovtID,
		&i.Email,
		&i.WalletAddress,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET firstname = $1, lastname = $2, govt_id = $3
WHERE email = $4
RETURNING  id, firstname, lastname, govt_id, email, wallet_address, created_at, updated_at
`

type UpdateUserParams struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	GovtID    string `json:"govt_id"`
	Email     string `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Firstname,
		arg.Lastname,
		arg.GovtID,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.GovtID,
		&i.Email,
		&i.WalletAddress,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
