package redis

import (
	"os"
	"strings"
)

func RedisConfig() string {
	return strings.TrimSpace(os.Getenv("REDIS_HOST"))
}
