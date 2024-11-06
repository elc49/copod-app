package repository

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/sirupsen/logrus"
)

type Upload struct {
	sql *sql.Queries
	log *logrus.Logger
}

func (r *Upload) Init(sql *sql.Queries) {
	r.sql = sql
	r.log = logger.GetLogger()
}

func (r *Upload) CreateUpload(ctx context.Context, args sql.CreateUploadParams) (*model.Upload, error) {
	u, err := r.sql.CreateUpload(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"param": args}).Errorf("repository: CreateUpload")
		return nil, err
	}

	return &model.Upload{
		ID:        u.ID,
		URI:       u.Uri,
		Type:      model.Doc(u.Type),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (r *Upload) GetUpload(ctx context.Context, args sql.GetUploadParams) (*model.Upload, error) {
	u, err := r.sql.GetUpload(ctx, args)
	if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"param": args}).Errorf("repository: GetUpload")
		return nil, err
	}

	return &model.Upload{
		ID:        u.ID,
		URI:       u.Uri,
		Type:      model.Doc(u.Type),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
