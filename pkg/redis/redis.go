package redis

import (
	"context"
	"strings"
	"time"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/pkg/logger"
	"github.com/redis/go-redis/v9"
)

// InitRedisClient 初始化 Redis 客户端
func InitRedisClient(cfg *config.RedisConfig) error {
	if cfg.RedisConnString == "" {
		RedisEnabled = false
		logger.SysLog("REDIS_CONN_STRING not set, Redis is not enabled")
		return nil
	}
	if cfg.SyncFrequency == "" {
		RedisEnabled = false
		logger.SysLog("SYNC_FREQUENCY not set, Redis is disabled")
		return nil
	}

	url := cfg.RedisConnString
	logger.SysLog("Redis is enabled")

	if cfg.RedisMasterName == "" {
		// Standalone mode
		logger.SysLog("Redis standalone mode")
		opt, err := redis.ParseURL(url)
		if err != nil {
			logger.SysFatal("failed to parse Redis connection string: " + err.Error())
			return err
		}
		RDB = redis.NewClient(opt)
	} else {
		// Cluster mode
		logger.SysLog("Redis cluster mode enabled")
		addrs := strings.Split(url, ",")
		RDB = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:      addrs,
			Password:   cfg.RedisPassword,
			MasterName: cfg.RedisMasterName,
			DB:         cfg.Database,
		})
	}

	// Test Redis connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		logger.SysFatal("Redis ping test failed: " + err.Error())
		return err
	}
	return nil
}

// RedisSet 设置 Redis 键值对
func Set(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	return RDB.Set(ctx, key, value, expiration).Err()
}

// RedisGet 获取 Redis 键的值
func Get(key string) (string, error) {
	ctx := context.Background()
	result, err := RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Handle the case where the key does not exist
	}
	return result, err
}

// RedisDel 删除 Redis 键
func Del(key string) error {
	ctx := context.Background()
	return RDB.Del(ctx, key).Err()
}

// RedisDecrease Redis 键值递减
func Decrease(key string, value int64) error {
	ctx := context.Background()
	return RDB.DecrBy(ctx, key, value).Err()
}
