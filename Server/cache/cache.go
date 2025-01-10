package cache

import (
	"context"
	"time"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/logger"
	"github.com/elc49/copod/util"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var c *cacheClient

type Cache interface {
	Get(ctx context.Context, key string, returnValue interface{}) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, expires time.Duration) error
	Redis() *redis.Client
}

type cacheClient struct {
	log *logrus.Logger
	rdb *redis.Client
}

func New() {
	log := logger.GetLogger()
	opt, err := redis.ParseURL(config.C.Database.Redis)
	if err != nil {
		log.WithError(err).Fatalln("cache: redis.ParseURL")
	}

	rdb := redis.NewClient(opt)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.WithError(err).Fatalln("cache: redis.Ping")
	}

	c = &cacheClient{log, rdb}
}

func GetCache() Cache {
	return c
}

func (c *cacheClient) Redis() *redis.Client {
	return c.rdb
}

func (c *cacheClient) Get(ctx context.Context, key string, returnValue interface{}) (interface{}, error) {
	result, err := c.rdb.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		return nil, nil
	case err != nil:
		c.log.WithError(err).WithFields(logrus.Fields{"key": key}).Errorf("cache: rdb.Get")
		return nil, err
	}

	if err := util.DecodeJson([]byte(result), returnValue); err != nil {
		c.log.WithError(err).Errorf("cache: DecodeJson")
		return nil, err
	}

	return returnValue, nil
}

func (c *cacheClient) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	b, err := util.EncodeJson(value)
	if err != nil {
		c.log.WithError(err).WithFields(logrus.Fields{"key": key, "value": value}).Errorf("cache: EncodeJson")
		return err
	}

	return c.rdb.Set(ctx, key, b, exp).Err()
}
