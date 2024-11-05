package config

import (
	"os"
	"strings"

	"github.com/elc49/copod/config/postgres"
	"github.com/elc49/copod/logger"
	"github.com/joho/godotenv"
)

var (
	C *config
)

type config struct {
	Tigris   Tigris
	Server   Server
	Database Database
}

func env() {
	godotenv.Load()
}

func New() {
	env()
	log := logger.GetLogger()
	c := config{}
	log.Infoln("Collecting configurations...")

	c.Server = serverConfig()
	c.Tigris = tigrisConfig()
	c.Database = databaseConfig()

	C = &c
	log.Infoln("Configurations...OK")
}

func getEnv() string {
	env := strings.TrimSpace(os.Getenv("ENV"))
	if env == "" {
		return "dev"
	}

	return env
}

func serverConfig() Server {
	var config Server

	config.Port = strings.TrimSpace(os.Getenv("PORT"))
	config.Env = getEnv()

	return config
}

func tigrisConfig() Tigris {
	var config Tigris

	config.SecretAccessKey = strings.TrimSpace(os.Getenv("AWS_SECRET_ACCESS_KEY"))
	config.AccessKeyId = strings.TrimSpace(os.Getenv("AWS_ACCESS_KEY_ID"))
	config.S3Endpoint = strings.TrimSpace(os.Getenv("AWS_S3_ENDPOINT_URL"))
	config.BucketName = strings.TrimSpace(os.Getenv("AWS_BUCKET_NAME"))
	config.Region = strings.TrimSpace(os.Getenv("AWS_S3_REGION"))

	return config
}

func databaseConfig() Database {
	var config Database

	config.Rdbms = postgres.PostgresConfig()

	return config
}

func IsProd() bool {
	return C != nil && (C.Server.Env == "staging" || C.Server.Env == "prod")
}
