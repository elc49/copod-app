package graph

import "github.com/elc49/copod/paystack"

//go:generate go run github.com/99designs/gqlgen generate --verbose

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	paystack paystack.Paystack
}

func New() Config {
	r := &Resolver{
		paystack.GetPaystackService(),
	}
	return Config{Resolvers: r}
}
