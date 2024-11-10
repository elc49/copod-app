package config

import (
	"os"
	"strings"
)

type Redis struct {
	Host     string
	Password string
	Db       string
}

func RedisConfig() Redis {
	var config Redis

	config.Host = strings.TrimSpace(os.Getenv("REDIS_ENDPOINT"))
	config.Password = strings.TrimSpace(os.Getenv("REDIS_PASSWORD"))
	config.Db = strings.TrimSpace(os.Getenv("REDIS_DATABASE"))

	return config
}
