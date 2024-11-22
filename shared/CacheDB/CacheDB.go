package CacheDB

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisConfig Redis配置结构
type RedisConfig struct {
	Addr         string
	Password     string
	DB           int
	PoolSize     int           // 连接池大小
	MinIdleConns int           // 最小空闲连接数
	DefaultTTL   time.Duration // 默认过期时间
}

// RedisDB 接口定义
type RedisDB interface {
	Get(key string) (string, error)
	GetWithContext(ctx context.Context, key string) (string, error)
	Set(key string, value string) error
	SetWithTTL(key string, value string, expiration time.Duration) error
	SetNX(key string, value string, expiration time.Duration) (bool, error)
	Del(key ...string) error
	Expire(key string, expiration time.Duration) error
	Exists(key string) (bool, error)
	Close() error
	Pipeline() redis.Pipeliner
}

// DefaultConfig 默认配置
func DefaultConfig() *RedisConfig {
	return &RedisConfig{
		Addr:         "redis:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
		DefaultTTL:   24 * time.Hour, // 默认24小时过期
	}
}

// NewRedisDB 创建Redis客户端
func NewRedisDB(config *RedisConfig) (RedisDB, error) {
	if config == nil {
		config = DefaultConfig()
	}

	client := redis.NewClient(&redis.Options{
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis连接失败: %v", err)
	}

	return &RedisDBImpl{
		Client:     client,
		DefaultTTL: config.DefaultTTL,
	}, nil
}

type RedisDBImpl struct {
	Client     *redis.Client
	DefaultTTL time.Duration
}

func (r *RedisDBImpl) Get(key string) (string, error) {
	return r.GetWithContext(context.Background(), key)
}

func (r *RedisDBImpl) GetWithContext(ctx context.Context, key string) (string, error) {
	result, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key不存在: %s", key)
	}
	return result, err
}

func (r *RedisDBImpl) Set(key string, value string) error {
	return r.SetWithTTL(key, value, r.DefaultTTL)
}

func (r *RedisDBImpl) SetWithTTL(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisDBImpl) SetNX(key string, value string, expiration time.Duration) (bool, error) {
	ctx := context.Background()
	return r.Client.SetNX(ctx, key, value, expiration).Result()
}

func (r *RedisDBImpl) Del(keys ...string) error {
	ctx := context.Background()
	return r.Client.Del(ctx, keys...).Err()
}

func (r *RedisDBImpl) Expire(key string, expiration time.Duration) error {
	ctx := context.Background()
	return r.Client.Expire(ctx, key, expiration).Err()
}

func (r *RedisDBImpl) Exists(key string) (bool, error) {
	ctx := context.Background()
	n, err := r.Client.Exists(ctx, key).Result()
	return n > 0, err
}

func (r *RedisDBImpl) Pipeline() redis.Pipeliner {
	return r.Client.Pipeline()
}

func (r *RedisDBImpl) Close() error {
	return r.Client.Close()
}
