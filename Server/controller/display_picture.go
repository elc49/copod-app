package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/google/uuid"
)

var displayPictureController *DisplayPicture

type DisplayPictureController interface {
	CreateDisplayPicture(context.Context, sql.CreateDisplayPictureParams) (*model.DisplayPicture, error)
	GetDisplayPictureByID(context.Context, uuid.UUID) (*model.DisplayPicture, error)
	UpdateDisplayPictureByID(context.Context, sql.UpdateDisplayPictureByIDParams) (*model.DisplayPicture, error)
	UpdateDisplayPictureVerificationByID(context.Context, sql.UpdateDisplayPictureVerificationByIDParams) (*model.DisplayPicture, error)
}

type DisplayPicture struct {
	r *repository.DisplayPicture
}

func (c *DisplayPicture) Init(sql *sql.Queries) {
	r := &repository.DisplayPicture{}
	r.Init(sql)
	c.r = r
	displayPictureController = c
}

func GetDisplayPictureController() DisplayPictureController {
	return displayPictureController
}

func (c *DisplayPicture) CreateDisplayPicture(ctx context.Context, args sql.CreateDisplayPictureParams) (*model.DisplayPicture, error) {
	return c.r.CreateDisplayPicture(ctx, args)
}

func (c *DisplayPicture) GetDisplayPictureByID(ctx context.Context, id uuid.UUID) (*model.DisplayPicture, error) {
	return c.r.GetDisplayPictureByID(ctx, id)
}

func (c *DisplayPicture) UpdateDisplayPictureByID(ctx context.Context, args sql.UpdateDisplayPictureByIDParams) (*model.DisplayPicture, error) {
	return c.r.UpdateDisplayPictureByID(ctx, args)
}

func (c *DisplayPicture) UpdateDisplayPictureVerificationByID(ctx context.Context, args sql.UpdateDisplayPictureVerificationByIDParams) (*model.DisplayPicture, error) {
	return c.r.UpdateDisplayPictureVerificationByID(ctx, args)
}
