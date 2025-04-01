package redis

import (
	"context"
	"strings"
	"time"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/pkg/logger"
	"github.com/redis/go-redis/v9"
)

var RedisEnabled = true
var RDB redis.UniversalClient // Changed to interface type

func InitRedisClient() (err error) {
	cfg := config.GetRedisConfig()
	if cfg.RedisConnString == "" {
		RedisEnabled = false
		logger.SysLog("REDIS_CONN_STRING not set, Redis is not enabled")
		return nil
	}
	if cfg.SyncFrequency == "" { // Corrected: Check for empty string
		RedisEnabled = false
		logger.SysLog("SYNC_FREQUENCY not set, Redis is disabled")
		return nil
	}

	url := cfg.RedisConnString
	logger.SysLog("Redis is enabled") // Moved log message

	if cfg.RedisMasterName == "" {
		// Standalone mode
		logger.SysLog("Redis standalone mode")
		opt, err := redis.ParseURL(url)
		if err != nil {
			logger.FatalLog("failed to parse Redis connection string: " + err.Error())
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
			DB:         cfg.Database, // Added: Use Database from config
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		logger.FatalLog("Redis ping test failed: " + err.Error())
		return err
	}
	return nil
}

// ParseRedisOption function is redundant and can be removed.
// The logic is now within InitRedisClient.

func RedisSet(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	return RDB.Set(ctx, key, value, expiration).Err()
}

func RedisGet(key string) (string, error) {
	ctx := context.Background()
	result, err := RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Handle the case where the key does not exist
	}
	return result, err
}

func RedisDel(key string) error {
	ctx := context.Background()
	return RDB.Del(ctx, key).Err()
}

func RedisDecrease(key string, value int64) error {
	ctx := context.Background()
	return RDB.DecrBy(ctx, key, value).Err()
}
