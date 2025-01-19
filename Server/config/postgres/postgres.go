package postgres

import (
	"os"
	"strconv"
	"strings"
)

type Postgres struct {
	DbName      string
	DbUser      string
	DbDriver    string
	DbHost      string
	DbPassword  string
	DbMigrate   bool
	DbMigration string
	Ssl         string
	Port        string
}

func PostgresConfig() Postgres {
	var config Postgres

	config.DbName = strings.TrimSpace(os.Getenv("POSTGRES_NAME"))
	config.DbUser = strings.TrimSpace(os.Getenv("POSTGRES_USER"))
	config.DbDriver = strings.TrimSpace(os.Getenv("POSTGRES_DRIVER"))
	config.DbHost = strings.TrimSpace(os.Getenv("POSTGRES_HOST"))
	config.DbPassword = strings.TrimSpace(os.Getenv("POSTGRES_PASSWORD"))
	forceMigrate, err := strconv.ParseBool(strings.TrimSpace(os.Getenv("POSTGRES_MIGRATE")))
	if err != nil {
		panic(err)
	}
	config.DbMigrate = forceMigrate
	config.DbMigration = strings.TrimSpace(os.Getenv("POSTGRES_MIGRATION"))
	config.Ssl = strings.TrimSpace(os.Getenv("POSTGRES_SSL"))
	config.Port = strings.TrimSpace(os.Getenv("POSTGRES_PORT"))

	return config
}
