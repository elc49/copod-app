package test

import (
	"context"

	"github.com/elc49/copod/config/postgres"
	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/sql"
	sqlc "github.com/elc49/copod/sql/sqlc"
	"github.com/sirupsen/logrus"
)

var (
	email           = RandomEmailAddress()
	superUserGovtId = RandomGovtID()

	docUri    = "https://doc.io/title"
	q         *sqlc.Queries
	superUser *model.User
)

func init() {
	opt := postgres.Postgres{
		DbUser:      "postgres",
		DbName:      "test",
		DbPassword:  "demo1234",
		DbMigration: "file://../sql/migration",
		DbMigrate:   true,
		DbHost:      "127.0.0.1",
		DbDriver:    "postgres",
	}

	// Database
	q = sql.InitDB(opt)
	// Controller
	// User
	u := controller.User{}
	u.Init(q)
	// Super user
	uc := controller.GetUserController()
	user, err := uc.CreateUser(context.Background(), email)
	if err != nil {
		logrus.WithError(err).Fatalln("test: init: CreateUser")
	}
	superUser = user
	// Payment
	pc := controller.Payment{}
	pc.Init(q)
	// Title
	tc := controller.Title{}
	tc.Init(q)
	// SupportingDoc
	sc := controller.SupportingDoc{}
	sc.Init(q)
}
