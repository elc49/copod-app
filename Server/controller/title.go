package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
)

var titleController *Title

type TitleController interface {
	CreateTitle(context.Context, sql.CreateTitleParams) (*model.Title, error)
	UpdateTitleVerification(context.Context, sql.UpdateTitleVerificationParams) (*model.Title, error)
}

type Title struct {
	r *repository.Title
}

func GetTitleController() TitleController {
	return titleController
}

func (c *Title) Init(sql *sql.Queries) {
	r := &repository.Title{}
	r.Init(sql)
	c.r = r
	titleController = c
}

func (c *Title) CreateTitle(ctx context.Context, args sql.CreateTitleParams) (*model.Title, error) {
	return c.r.CreateTitle(ctx, args)
}

func (c *Title) UpdateTitleVerification(ctx context.Context, args sql.UpdateTitleVerificationParams) (*model.Title, error) {
	return c.r.UpdateTitleVerification(ctx, args)
}
