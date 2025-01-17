package sql

import (
	"context"
	"fmt"

	db "database/sql"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/config/postgres"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/elc49/copod/tigris"
	"github.com/golang-migrate/migrate/v4"
	p "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func InitDB(opt postgres.Postgres) *sql.Queries {
	log := logger.GetLogger()
	uri := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s", opt.DbUser, opt.DbPassword, opt.DbHost, opt.DbName, opt.Ssl)

	conn, _ := db.Open(opt.DbDriver, uri)

	if err := conn.Ping(); err != nil {
		log.WithError(err).Fatalln("postgres: Ping")
		return nil
	} else {
		log.Infoln("Database connection...OK")
	}

	conn.Exec(fmt.Sprintf("CREATE EXTENSION IF NOT EXISTS %q;", "uuid-ossp"))
	conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis;")
	conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis_rasters; --OPTIONAL")
	conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis_topology; --OPTIONAL")

	if err := runMigration(opt, conn, log); err != nil {
		log.WithError(err).Fatalln("sql: runMigration")
	} else {
		log.Infoln("Write table schema...OK")
	}

	return sql.New(conn)
}

func runMigration(opt postgres.Postgres, conn *db.DB, log *logrus.Logger) error {
	d, err := p.WithInstance(conn, &p.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(opt.DbMigration, opt.DbDriver, d)
	if err != nil {
		return err
	}

	if opt.DbMigrate {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			return err
		}

		if config.C != nil {
			if err := tigris.T.DeleteObjects(context.Background()); err != nil {
				log.WithError(err).Fatalln("sql: DeleteObjects")
				return err
			}
		}
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
