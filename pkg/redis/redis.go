package initialize

import (
	"context"
	"strings"
	"time"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/pkg/logger"
	"github.com/redis/go-redis/v9"
)

var RedisEnabled = true
var RDB *redis.UniversalOptions

func InitRedisClient(cfg *config.RedisConfig) (err error) {
	if cfg.RedisConnString == "" {
		RedisEnabled = false
		logger.SysLog("REDIS_CONN_STRING not set, Redis is not enabled")
		return nil
	}
	if cfg.SyncFrequency == 0 {
		RedisEnabled = false
		logger.SysLog("SYNC_FREQUENCY not set, Redis is disabled")
		return nil
	}
	url := cfg.RedisConnString
	if cfg.RedisMasterName == "" {
		logger.SysLog("Redis is enabled")
		opt, err := redis.ParseURL(url)
		if err != nil {
			logger.FatalLog("failed to parse Redis connection string: " + err.Error())
		}
		RDB = redis.NewClient(opt)
	} else {
		// cluster mode
		logger.SysLog("Redis cluster mode enabled")
		RDB = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:      strings.Split(url, ","),
			Password:   cfg.RedisPassword,
			MasterName: cfg.RedisMasterName,
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		logger.FatalLog("Redis ping test failed: " + err.Error())
	}
	return err
}

func ParseRedisOption(url string) *redis.Options {
	opt, err := redis.ParseURL(url)
	if err != nil {
		logger.FatalLog("failed to parse Redis connection string: " + err.Error())
	}
	return opt
}
func RedisSet(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	return RDB.Set(ctx, key, value, expiration).Err()
}

func RedisGet(key string) (string, error) {
	ctx := context.Background()
	return RDB.Get(ctx, key).Result()
}

func RedisDel(key string) error {
	ctx := context.Background()
	return RDB.Del(ctx, key).Err()
}

func RedisDecrease(key string, value int64) error {
	ctx := context.Background()
	return RDB.DecrBy(ctx, key, value).Err()
}
