// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: upload.sql

package sql

import (
	"context"
	"database/sql"
)

const createUpload = `-- name: CreateUpload :one
INSERT INTO uploads (
  type, uri, verification, wallet_address
) VALUES (
  $1, $2, $3, $4
) RETURNING id, type, uri, verification, land_id, wallet_address, created_at, updated_at
`

type CreateUploadParams struct {
	Type          string         `json:"type"`
	Uri           string         `json:"uri"`
	Verification  string         `json:"verification"`
	WalletAddress sql.NullString `json:"wallet_address"`
}

func (q *Queries) CreateUpload(ctx context.Context, arg CreateUploadParams) (Upload, error) {
	row := q.db.QueryRowContext(ctx, createUpload,
		arg.Type,
		arg.Uri,
		arg.Verification,
		arg.WalletAddress,
	)
	var i Upload
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Uri,
		&i.Verification,
		&i.LandID,
		&i.WalletAddress,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
