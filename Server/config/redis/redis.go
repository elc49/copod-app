package redis

import (
	"os"
	"strconv"
	"strings"
)

type Redis struct {
	Host     string
	Password string
	Db       int
}

func RedisConfig() Redis {
	var config Redis

	config.Host = strings.TrimSpace(os.Getenv("REDIS_ENDPOINT"))
	config.Password = strings.TrimSpace(os.Getenv("REDIS_PASSWORD"))
	db, err := strconv.Atoi(strings.TrimSpace(os.Getenv("REDIS_DATABASE")))
	if err != nil {
		panic(err)
	}
	config.Db = db

	return config
}
