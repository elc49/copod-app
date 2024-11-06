package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
)

var PC *Upload

type UploadController interface {
	CreateUpload(context.Context, sql.CreateUploadParams) (*model.Upload, error)
	GetUpload(context.Context, sql.GetUploadParams) (*model.Upload, error)
}

var _ UploadController = (*Upload)(nil)

type Upload struct {
	r *repository.Upload
}

func (c *Upload) Init(sql *sql.Queries) {
	r := &repository.Upload{}
	r.Init(sql)
	c.r = r
	PC = c
}

func (c *Upload) CreateUpload(ctx context.Context, args sql.CreateUploadParams) (*model.Upload, error) {
	return c.r.CreateUpload(ctx, args)
}

func (c *Upload) GetUpload(ctx context.Context, args sql.GetUploadParams) (*model.Upload, error) {
	return c.r.GetUpload(ctx, args)
}
