package sql

import (
	"fmt"

	db "database/sql"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/config/postgres"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/golang-migrate/migrate/v4"
	gm "github.com/golang-migrate/migrate/v4/database/postgres"
)

func InitDB(opt postgres.Postgres) *sql.Queries {
	log := logger.GetLogger()
	uri := fmt.Sprintf("user=%s password=%s host=%s dbname=%s", opt.DbUser, opt.DbPassword, opt.DbHost, opt.DbName)
	if !config.IsProd() {
		uri += " sslmode=disable"
	}

	conn, _ := db.Open(opt.DbDriver, uri)

	if err := conn.Ping(); err != nil {
		log.WithError(err).Fatalln("postgres: Ping")
		return nil
	} else {
		log.Infoln("Database connection...OK")
	}

	conn.Exec(fmt.Sprintf("CREATE EXTENSION IF NOT EXISTS %q;", "uuid-ossp"))
	conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis")
	conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis_rasters; --OPTIONAL")
	conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis_topology; --OPTIONAL")

	if err := runMigration(opt.DbMigration, opt.DbMigrate, conn); err != nil {
		log.WithError(err).Fatalln("sql: runMigration")
	} else {
		log.Infoln("Write table schema...OK")
	}

	return sql.New(conn)
}

func runMigration(migration string, forceMigrate bool, conn *db.DB) error {
	driver, err := gm.WithInstance(conn, &gm.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migration, "postgres", driver)
	if err != nil {
		return nil
	}

	if forceMigrate {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			return nil
		}
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil
	}

	return nil
}
