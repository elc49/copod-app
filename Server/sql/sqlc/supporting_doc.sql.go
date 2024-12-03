// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: supporting_doc.sql

package sql

import (
	"context"
)

const createSupportDoc = `-- name: CreateSupportDoc :one
INSERT INTO support_docs (
  email, govt_id
) VALUES (
  $1, $2
) RETURNING id, govt_id, verification, email, created_at, updated_at
`

type CreateSupportDocParams struct {
	Email  string `json:"email"`
	GovtID string `json:"govt_id"`
}

func (q *Queries) CreateSupportDoc(ctx context.Context, arg CreateSupportDocParams) (SupportDoc, error) {
	row := q.db.QueryRowContext(ctx, createSupportDoc, arg.Email, arg.GovtID)
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

const getSupportDocByEmail = `-- name: GetSupportDocByEmail :one
SELECT id, govt_id, verification, email, created_at, updated_at FROM support_docs
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetSupportDocByEmail(ctx context.Context, email string) (SupportDoc, error) {
	row := q.db.QueryRowContext(ctx, getSupportDocByEmail, email)
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

const updateSupportDocByEmail = `-- name: UpdateSupportDocByEmail :one
UPDATE support_docs SET govt_id = $1, verification = $2
WHERE email = $3
RETURNING id, govt_id, verification, email, created_at, updated_at
`

type UpdateSupportDocByEmailParams struct {
	GovtID       string `json:"govt_id"`
	Verification string `json:"verification"`
	Email        string `json:"email"`
}

func (q *Queries) UpdateSupportDocByEmail(ctx context.Context, arg UpdateSupportDocByEmailParams) (SupportDoc, error) {
	row := q.db.QueryRowContext(ctx, updateSupportDocByEmail, arg.GovtID, arg.Verification, arg.Email)
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
