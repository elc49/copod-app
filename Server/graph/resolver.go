package graph

import (
	"github.com/elc49/copod/cache"
	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/paystack"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

//go:generate go run github.com/99designs/gqlgen generate --verbose

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	paystack                 paystack.Paystack
	redis                    *redis.Client
	log                      *logrus.Logger
	titleController          controller.TitleController
	supportDocController     controller.SupportingDocController
	paymentController        controller.PaymentController
	userController           controller.UserController
	onboardingController     controller.OnboardingController
	displayPictureController controller.DisplayPictureController
}

func New() Config {
	r := &Resolver{
		paystack.GetPaystackService(),
		cache.GetCache().Redis(),
		logger.GetLogger(),
		controller.GetTitleController(),
		controller.GetSupportingDocController(),
		controller.GetPaymentController(),
		controller.GetUserController(),
		controller.GetOnboardingController(),
		controller.GetDisplayPictureController(),
	}
	return Config{Resolvers: r}
}
