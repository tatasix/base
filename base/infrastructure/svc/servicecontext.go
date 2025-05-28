package svc

import (
	"base/infrastructure/config"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger := logx.WithContext(context.Background())
	redisConf := redis.RedisConf{
		Host: c.Redis.Host,
		Type: c.Redis.Type,
		Pass: c.Redis.Pass,
		Tls:  c.Redis.Tls,
	}

	redisClient, err := redis.NewRedis(redisConf)
	if err != nil {
		logger.Errorf("Failed to create Redis client: %v", err)
	}
	return &ServiceContext{
		Config: c,
		Redis:  redisClient,
	}
}
