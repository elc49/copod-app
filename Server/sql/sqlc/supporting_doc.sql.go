// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: supporting_doc.sql

package sql

import (
	"context"

	"github.com/google/uuid"
)

const createSupportDoc = `-- name: CreateSupportDoc :one
INSERT INTO support_docs (
  email, url
) VALUES (
  $1, $2
) RETURNING id, url, verification, email, created_at, updated_at
`

type CreateSupportDocParams struct {
	Email string `json:"email"`
	Url   string `json:"url"`
}

func (q *Queries) CreateSupportDoc(ctx context.Context, arg CreateSupportDocParams) (SupportDoc, error) {
	row := q.db.QueryRowContext(ctx, createSupportDoc, arg.Email, arg.Url)
	var i SupportDoc
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Verification,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSupportDocByEmail = `-- name: GetSupportDocByEmail :one
SELECT id, url, verification, email, created_at, updated_at FROM support_docs
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetSupportDocByEmail(ctx context.Context, email string) (SupportDoc, error) {
	row := q.db.QueryRowContext(ctx, getSupportDocByEmail, email)
	var i SupportDoc
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Verification,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSupportDocByID = `-- name: GetSupportDocByID :one
SELECT id, url, verification, email, created_at, updated_at FROM support_docs
WHERE id = $1
`

func (q *Queries) GetSupportDocByID(ctx context.Context, id uuid.UUID) (SupportDoc, error) {
	row := q.db.QueryRowContext(ctx, getSupportDocByID, id)
	var i SupportDoc
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Verification,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSupportingDocsByVerification = `-- name: GetSupportingDocsByVerification :many
SELECT id, url, verification, email, created_at, updated_at FROM support_docs
WHERE verification = $1
`

func (q *Queries) GetSupportingDocsByVerification(ctx context.Context, verification string) ([]SupportDoc, error) {
	rows, err := q.db.QueryContext(ctx, getSupportingDocsByVerification, verification)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SupportDoc{}
	for rows.Next() {
		var i SupportDoc
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.Verification,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSupportDocByID = `-- name: UpdateSupportDocByID :one
UPDATE support_docs SET url = $1, verification = $2
WHERE id = $3
RETURNING id, url, verification, email, created_at, updated_at
`

type UpdateSupportDocByIDParams struct {
	Url          string    `json:"url"`
	Verification string    `json:"verification"`
	ID           uuid.UUID `json:"id"`
}

func (q *Queries) UpdateSupportDocByID(ctx context.Context, arg UpdateSupportDocByIDParams) (SupportDoc, error) {
	row := q.db.QueryRowContext(ctx, updateSupportDocByID, arg.Url, arg.Verification, arg.ID)
	var i SupportDoc
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Verification,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSupportDocVerificationByID = `-- name: UpdateSupportDocVerificationByID :one
UPDATE support_docs SET verification = $1
WHERE id = $2
RETURNING id, url, verification, email, created_at, updated_at
`

type UpdateSupportDocVerificationByIDParams struct {
	Verification string    `json:"verification"`
	ID           uuid.UUID `json:"id"`
}

func (q *Queries) UpdateSupportDocVerificationByID(ctx context.Context, arg UpdateSupportDocVerificationByIDParams) (SupportDoc, error) {
	row := q.db.QueryRowContext(ctx, updateSupportDocVerificationByID, arg.Verification, arg.ID)
	var i SupportDoc
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Verification,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
