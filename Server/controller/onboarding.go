package controller

import (
	"context"

	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/repository"
	sql "github.com/elc49/copod/sql/sqlc"
)

var onboardingController *Onboarding

type OnboardingController interface {
	CreateOnboarding(context.Context, sql.CreateOnboardingParams) (*model.Onboarding, error)
}

type Onboarding struct {
	r *repository.Onboarding
}

func (o *Onboarding) Init(sql *sql.Queries) {
	r := &repository.Onboarding{}
	r.Init(sql)
	o.r = r
	onboardingController = o
}

func GetOnboardingController() OnboardingController {
	return onboardingController
}

func (o *Onboarding) CreateOnboarding(ctx context.Context, args sql.CreateOnboardingParams) (*model.Onboarding, error) {
	return o.r.CreateOnboarding(ctx, args)
}
