package test

import (
	"github.com/elc49/copod/config/postgres"
	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/sql"
)

var (
	email         = "email@exmpl.com"
	walletAddress = "0x41eD3Ce6DC13fD4F67Eb715f5c3B105Bc7FA8D45"
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
	q := sql.InitDB(opt)
	// Controller
	// User
	u := controller.User{}
	u.Init(q)
}
