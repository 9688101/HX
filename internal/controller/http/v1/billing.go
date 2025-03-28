package v1

import (
	"github.com/9688101/HX/common/config"
	"github.com/9688101/HX/common/ctxkey"
	"github.com/9688101/HX/internal/entity"
	relaymodel "github.com/9688101/HX/relay/model"
	"github.com/gin-gonic/gin"
)

func GetSubscription(c *gin.Context) {
	var remainQuota int64
	var usedQuota int64
	var err error
	var token *entity.Token
	var expiredTime int64
	if config.DisplayTokenStatEnabled {
		tokenId := c.GetInt(ctxkey.TokenId)
		token, err = entity.GetTokenById(tokenId)
		if err == nil {
			expiredTime = token.ExpiredTime
			remainQuota = token.RemainQuota
			usedQuota = token.UsedQuota
		}
	} else {
		userId := c.GetInt(ctxkey.Id)
		remainQuota, err = entity.GetUserQuota(userId)
		if err != nil {
			usedQuota, err = entity.GetUserUsedQuota(userId)
		}
	}
	if expiredTime <= 0 {
		expiredTime = 0
	}
	if err != nil {
		Error := relaymodel.Error{
			Message: err.Error(),
			Type:    "upstream_error",
		}
		c.JSON(200, gin.H{
			"error": Error,
		})
		return
	}
	quota := remainQuota + usedQuota
	amount := float64(quota)
	if config.DisplayInCurrencyEnabled {
		amount /= config.QuotaPerUnit
	}
	if token != nil && token.UnlimitedQuota {
		amount = 100000000
	}
	subscription := OpenAISubscriptionResponse{
		Object:             "billing_subscription",
		HasPaymentMethod:   true,
		SoftLimitUSD:       amount,
		HardLimitUSD:       amount,
		SystemHardLimitUSD: amount,
		AccessUntil:        expiredTime,
	}
	c.JSON(200, subscription)
	return
}

func GetUsage(c *gin.Context) {
	var quota int64
	var err error
	var token *entity.Token
	if config.DisplayTokenStatEnabled {
		tokenId := c.GetInt(ctxkey.TokenId)
		token, err = entity.GetTokenById(tokenId)
		quota = token.UsedQuota
	} else {
		userId := c.GetInt(ctxkey.Id)
		quota, err = entity.GetUserUsedQuota(userId)
	}
	if err != nil {
		Error := relaymodel.Error{
			Message: err.Error(),
			Type:    "one_api_error",
		}
		c.JSON(200, gin.H{
			"error": Error,
		})
		return
	}
	amount := float64(quota)
	if config.DisplayInCurrencyEnabled {
		amount /= config.QuotaPerUnit
	}
	usage := OpenAIUsageResponse{
		Object:     "list",
		TotalUsage: amount * 100,
	}
	c.JSON(200, usage)
	return
}
