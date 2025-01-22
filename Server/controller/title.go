package controller

import (
	"context"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/ethereum"
	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/google/uuid"
)

var titleController *Title

type TitleController interface {
	CreateTitle(context.Context, sql.CreateTitleParams) (*model.Title, error)
	UpdateTitleByID(context.Context, sql.UpdateTitleByIDParams) (*model.Title, error)
	GetTitlesByEmailAndVerification(context.Context, sql.GetTitlesByEmailAndVerificationParams) ([]*model.Title, error)
	GetTitleByID(context.Context, uuid.UUID) (*model.Title, error)
	UpdateTitleVerificationByID(context.Context, sql.UpdateTitleVerificationByIDParams, ethereum.LandDetails) (*model.Title, error)
}

type Title struct {
	r          *repository.Title
	ethBackend ethereum.EthBackend
}

func GetTitleController() TitleController {
	return titleController
}

func (c *Title) Init(sql *sql.Queries) {
	r := &repository.Title{}
	r.Init(sql)
	c.r = r
	c.ethBackend = ethereum.GetEthBackend()
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

func (c *Title) GetTitleByID(ctx context.Context, id uuid.UUID) (*model.Title, error) {
	return c.r.GetTitleByID(ctx, id)
}

func (c *Title) UpdateTitleVerificationByID(ctx context.Context, args sql.UpdateTitleVerificationByIDParams, landDetails ethereum.LandDetails) (*model.Title, error) {
	// Avoid running ethereum in test env
	if args.Verification == model.VerificationVerified.String() && !config.IsTest() {
		if err := c.ethBackend.RegisterLand(context.Background(), landDetails); err != nil {
			return nil, err
		}
	}

	return c.r.UpdateTitleVerificationByID(ctx, args)
}
