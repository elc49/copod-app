package repository

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/sirupsen/logrus"
)

type Title struct {
	sql *sql.Queries
	log *logrus.Logger
}

func (r *Title) Init(sql *sql.Queries) {
	r.sql = sql
	r.log = logger.GetLogger()
}

func (r *Title) CreateTitle(ctx context.Context, args sql.CreateTitleParams) (*model.Title, error) {
	t, err := r.sql.CreateTitle(ctx, args)
	if err != nil {
		r.log.WithError(err).Errorf("repository: CreateTitle")
		return nil, err
	}

	return &model.Title{
		ID:        t.ID,
		Title:     t.Title,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, nil
}
