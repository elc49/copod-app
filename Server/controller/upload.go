package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
)

var uploadController *Upload

type UploadController interface {
	CreateUpload(context.Context, sql.CreateUploadParams) (*model.Upload, error)
}

var _ UploadController = (*Upload)(nil)

type Upload struct {
	r *repository.Upload
}

func (c *Upload) Init(sql *sql.Queries) {
	r := &repository.Upload{}
	r.Init(sql)
	c.r = r
	uploadController = c
}

func GetUploadController() UploadController {
	return uploadController
}

func (c *Upload) CreateUpload(ctx context.Context, args sql.CreateUploadParams) (*model.Upload, error) {
	return c.r.CreateUpload(ctx, args)
}
