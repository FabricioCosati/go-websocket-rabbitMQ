package redis

import (
	"encoding/json"
	"time"

	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/config"
	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Redis *redis.Client
}

func NewRedisClient(cfg config.Config) *RedisClient {
	client := redis.NewClient(
		&redis.Options{
			Addr: cfg.RedisConn,
		},
	)

	return &RedisClient{
		Redis: client,
	}
}

func (s *RedisClient) Set(ctx *gin.Context, key string, value interface{}) error {
	exp := time.Duration(30 * time.Minute)

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return s.Redis.Set(ctx, key, b, exp).Err()
}

func (s *RedisClient) Get(ctx *gin.Context, key string) (interface{}, error) {
	value, err := s.Redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return value, nil
}
