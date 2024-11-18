// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: payment.sql

package sql

import (
	"context"

	"github.com/google/uuid"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments (
  email, amount, currency, reason, status, reference_id, title_id, wallet_address
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id, email, amount, currency, reason, status, reference_id, wallet_address, title_id, created_at, updated_at
`

type CreatePaymentParams struct {
	Email         string        `json:"email"`
	Amount        int32         `json:"amount"`
	Currency      string        `json:"currency"`
	Reason        string        `json:"reason"`
	Status        string        `json:"status"`
	ReferenceID   string        `json:"reference_id"`
	TitleID       uuid.NullUUID `json:"title_id"`
	WalletAddress string        `json:"wallet_address"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, createPayment,
		arg.Email,
		arg.Amount,
		arg.Currency,
		arg.Reason,
		arg.Status,
		arg.ReferenceID,
		arg.TitleID,
		arg.WalletAddress,
	)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Amount,
		&i.Currency,
		&i.Reason,
		&i.Status,
		&i.ReferenceID,
		&i.WalletAddress,
		&i.TitleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPaymentByReferenceID = `-- name: GetPaymentByReferenceID :one
SELECT id, email, amount, currency, reason, status, reference_id, wallet_address, title_id, created_at, updated_at FROM payments
WHERE reference_id = $1
LIMIT 1
`

func (q *Queries) GetPaymentByReferenceID(ctx context.Context, referenceID string) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPaymentByReferenceID, referenceID)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Amount,
		&i.Currency,
		&i.Reason,
		&i.Status,
		&i.ReferenceID,
		&i.WalletAddress,
		&i.TitleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPaymentTitleByID = `-- name: GetPaymentTitleByID :one
SELECT id, title, verification, email, wallet_address, created_at, updated_at FROM titles
WHERE id = $1
`

func (q *Queries) GetPaymentTitleByID(ctx context.Context, id uuid.UUID) (Title, error) {
	row := q.db.QueryRowContext(ctx, getPaymentTitleByID, id)
	var i Title
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Verification,
		&i.Email,
		&i.WalletAddress,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPaymentsByStatus = `-- name: GetPaymentsByStatus :many
SELECT id, email, amount, currency, reason, status, reference_id, wallet_address, title_id, created_at, updated_at FROM payments
WHERE status = $1
`

func (q *Queries) GetPaymentsByStatus(ctx context.Context, status string) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentsByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Payment{}
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Amount,
			&i.Currency,
			&i.Reason,
			&i.Status,
			&i.ReferenceID,
			&i.WalletAddress,
			&i.TitleID,
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

const updatePaymentStatus = `-- name: UpdatePaymentStatus :one
UPDATE payments SET status = $1
WHERE reference_id = $2
RETURNING id, email, amount, currency, reason, status, reference_id, wallet_address, title_id, created_at, updated_at
`

type UpdatePaymentStatusParams struct {
	Status      string `json:"status"`
	ReferenceID string `json:"reference_id"`
}

func (q *Queries) UpdatePaymentStatus(ctx context.Context, arg UpdatePaymentStatusParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, updatePaymentStatus, arg.Status, arg.ReferenceID)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Amount,
		&i.Currency,
		&i.Reason,
		&i.Status,
		&i.ReferenceID,
		&i.WalletAddress,
		&i.TitleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
