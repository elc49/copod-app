package config

import (
	"os"
	"strings"

	"github.com/elc49/copod/config/postgres"
	"github.com/elc49/copod/config/redis"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	C *config
)

type config struct {
	Tigris   Tigris
	Server   Server
	Database Database
	Paystack Paystack
	Ipinfo   Ipinfo
	Ethereum Ethereum
	Resend   Resend
	Sentry   Sentry
}

func env() {
	godotenv.Load()
}

func New() {
	env()
	log := logrus.New()
	c := config{}
	log.Infoln("Collecting configurations...")

	c.Server = serverConfig()
	c.Tigris = tigrisConfig()
	c.Database = databaseConfig()
	c.Paystack = paystackConfig()
	c.Ipinfo = ipinfoConfig()
	c.Ethereum = ethereumConfig()
	c.Resend = resendConfig()
	c.Sentry = sentryConfig()

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
	config.Redis = redis.RedisConfig()

	return config
}

func IsProd() bool {
	return C != nil && (C.Server.Env == "staging" || C.Server.Env == "prod")
}

func paystackConfig() Paystack {
	var config Paystack

	config.BaseApi = strings.TrimSpace(os.Getenv("PAYSTACK_API"))
	config.SecretKey = strings.TrimSpace(os.Getenv("PAYSTACK_SECRET_KEY"))
	config.MobileTestAccount = strings.TrimSpace(os.Getenv("PAYSTACK_MOBILE_TEST_ACCOUNT"))
	config.LandFees = strings.TrimSpace(os.Getenv("PAYSTACK_LAND_FEES"))

	return config
}

func ipinfoConfig() Ipinfo {
	var config Ipinfo

	config.ApiKey = strings.TrimSpace(os.Getenv("IPINFO_API_KEY"))

	return config
}

func ethereumConfig() Ethereum {
	var config Ethereum

	config.InfuraApi = strings.TrimSpace(os.Getenv("INFURA_API"))
	config.RegistryContractAddress = strings.TrimSpace(os.Getenv("REGISTRY_CONTRACT_ADDRESS"))
	config.SigningAccountKey = strings.TrimSpace(os.Getenv("SIGNING_ACCOUNT_KEY"))

	return config
}

func resendConfig() Resend {
	var config Resend

	config.ApiKey = strings.TrimSpace(os.Getenv("RESEND_API_KEY"))

	return config
}

func sentryConfig() Sentry {
	var config Sentry

	config.Dsn = strings.TrimSpace(os.Getenv("SENTRY_DSN"))

	return config
}
