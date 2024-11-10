package config

import (
	"github.com/elc49/copod/config/postgres"
	"github.com/elc49/copod/config/redis"
)

type Database struct {
	Rdbms postgres.Postgres
	Redis redis.Redis
}
