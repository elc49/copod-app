package test

import (
	"github.com/elc49/copod/config/postgres"
	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/sql"
	sqlc "github.com/elc49/copod/sql/sqlc"
)

var (
	docUri = "https://doc.io/title"
	q      *sqlc.Queries
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
	// Upload
	p := controller.Upload{}
	p.Init(q)
}
