package CacheDB

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisDB interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Del(key string) error
	Expire(key string, expiration time.Duration) error
	Close() error
}

func NewRedisDB() RedisDB {
	return &RedisDBImpl{
		Client: redis.NewClient(&redis.Options{
			Addr:     "redis:6379", // Redis服务器地址
			Password: "",           // Redis服务器密码
			DB:       0,            // Redis数据库索引
		}),
	}
}

type RedisDBImpl struct {
	Client *redis.Client
}

func (r *RedisDBImpl) Get(key string) (string, error) {
	return r.Client.Get(context.Background(), key).Result()
}

func (r *RedisDBImpl) Set(key string, value string) error {
	return r.Client.Set(context.Background(), key, value, 0).Err()
}

func (r *RedisDBImpl) Del(key string) error {
	return r.Client.Del(context.Background(), key).Err()
}

func (r *RedisDBImpl) Expire(key string, expiration time.Duration) error {
	return r.Client.Expire(context.Background(), key, expiration).Err()
}

func (r *RedisDBImpl) Close() error {
	return r.Client.Close()
}
