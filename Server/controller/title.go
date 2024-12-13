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
	UpdateTitleByID(context.Context, sql.UpdateTitleByIDParams) (*model.Title, error)
	GetTitlesByEmailAndVerification(context.Context, sql.GetTitlesByEmailAndVerificationParams) ([]*model.Title, error)
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

func (c *Title) UpdateTitleByID(ctx context.Context, args sql.UpdateTitleByIDParams) (*model.Title, error) {
	return c.r.UpdateTitleByID(ctx, args)
}

func (c *Title) GetTitlesByEmailAndVerification(ctx context.Context, args sql.GetTitlesByEmailAndVerificationParams) ([]*model.Title, error) {
	return c.r.GetTitlesByEmailAndVerification(ctx, args)
}
