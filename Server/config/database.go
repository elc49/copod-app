package config

import "github.com/elc49/copod/config/postgres"

type Database struct {
	Rdbms postgres.Postgres
}
