package repository

import (
	"context"
	db "database/sql"

	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/sirupsen/logrus"
)

type EarlySignup struct {
	sql *sql.Queries
	log *logrus.Logger
}

func (r *EarlySignup) Init(sql *sql.Queries) {
	r.sql = sql
	r.log = logger.GetLogger()
}

func (r *EarlySignup) CreateEarlySignup(ctx context.Context, email string) (*string, error) {
	exists, err := r.sql.GetEarlySignupByEmail(ctx, email)
	if err != nil && err == db.ErrNoRows {
		e, err := r.sql.CreateEarlySignup(ctx, email)
		if err != nil {
			r.log.WithError(err).WithFields(logrus.Fields{"email": email}).Errorf("repository: CreateEarlySignup")
			return nil, err
		}
		return &e.Email, nil
	} else if err != nil {
		r.log.WithError(err).WithFields(logrus.Fields{"email": email}).Errorf("repository: GetEarlySignupByEmail")
		return nil, err
	}
	return &exists, nil
}
