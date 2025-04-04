package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/pkg/logger"
	"github.com/9688101/HX/pkg/redis"
	"github.com/9688101/HX/pkg/rl"
	"github.com/gin-gonic/gin"
)

var timeFormat = "2006-01-02T15:04:05.000Z"

var inMemoryRateLimiter rl.InMemoryRateLimiter

func redisRateLimiter(c *gin.Context, maxRequestNum int, duration int64, mark string) {
	ctx := context.Background()
	rdb := redis.RDB
	cfg := config.GetRateLimitConfig()
	key := "rateLimit:" + mark + c.ClientIP()
	listLength, err := rdb.LLen(ctx, key).Result()
	if err != nil {
		logger.Error(ctx, err.Error())
		c.Status(http.StatusInternalServerError)
		c.Abort()
		return
	}
	if listLength < int64(maxRequestNum) {
		rdb.LPush(ctx, key, time.Now().Format(timeFormat))
		rdb.Expire(ctx, key, cfg.RateLimitKeyExpirationDuration)
	} else {
		oldTimeStr, _ := rdb.LIndex(ctx, key, -1).Result()
		oldTime, err := time.Parse(timeFormat, oldTimeStr)
		if err != nil {
			logger.Error(ctx, err.Error())
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}
		nowTimeStr := time.Now().Format(timeFormat)
		nowTime, err := time.Parse(timeFormat, nowTimeStr)
		if err != nil {
			logger.Error(ctx, err.Error())
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}
		// time.Since will return negative number!
		// See: https://stackoverflow.com/questions/50970900/why-is-time-since-returning-negative-durations-on-windows
		if int64(nowTime.Sub(oldTime).Seconds()) < duration {
			rdb.Expire(ctx, key, cfg.RateLimitKeyExpirationDuration)
			c.Status(http.StatusTooManyRequests)
			c.Abort()
			return
		} else {
			rdb.LPush(ctx, key, time.Now().Format(timeFormat))
			rdb.LTrim(ctx, key, 0, int64(maxRequestNum-1))
			rdb.Expire(ctx, key, cfg.RateLimitKeyExpirationDuration)
		}
	}
}

func memoryRateLimiter(c *gin.Context, maxRequestNum int, duration int64, mark string) {
	key := mark + c.ClientIP()
	if !inMemoryRateLimiter.Request(key, maxRequestNum, duration) {
		c.Status(http.StatusTooManyRequests)
		c.Abort()
		return
	}
}

func rateLimitFactory(maxRequestNum int, duration int64, mark string) func(c *gin.Context) {
	if maxRequestNum == 0 || config.GetGeneralConfig().DebugEnabled {
		return func(c *gin.Context) {
			c.Next()
		}
	}
	if redis.RedisEnabled {
		return func(c *gin.Context) {
			redisRateLimiter(c, maxRequestNum, duration, mark)
		}
	} else {
		// It's safe to call multi times.
		inMemoryRateLimiter.Init(config.GetRateLimitConfig().RateLimitKeyExpirationDuration)
		return func(c *gin.Context) {
			memoryRateLimiter(c, maxRequestNum, duration, mark)
		}
	}
}

func GlobalWebRateLimit() func(c *gin.Context) {
	cfg := config.GetRateLimitConfig()
	return rateLimitFactory(cfg.GlobalWebRateLimitNum, cfg.GlobalWebRateLimitDuration, "GW")
}

func GlobalAPIRateLimit() func(c *gin.Context) {
	cfg := config.GetRateLimitConfig()

	return rateLimitFactory(cfg.GlobalApiRateLimitNum, cfg.GlobalApiRateLimitDuration, "GA")
}

func CriticalRateLimit() func(c *gin.Context) {
	cfg := config.GetRateLimitConfig()

	return rateLimitFactory(cfg.CriticalRateLimitNum, cfg.CriticalRateLimitDuration, "CT")
}

func DownloadRateLimit() func(c *gin.Context) {
	cfg := config.GetRateLimitConfig()

	return rateLimitFactory(cfg.DownloadRateLimitNum, cfg.DownloadRateLimitDuration, "DW")
}

func UploadRateLimit() func(c *gin.Context) {
	cfg := config.GetRateLimitConfig()

	return rateLimitFactory(cfg.UploadRateLimitNum, cfg.UploadRateLimitDuration, "UP")
}
