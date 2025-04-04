package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		fmt.Println(888, err)

		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Invalid request body",
		})
		return
	}

	// TODO: 从数据库验证用户名和密码
	// 这里使用示例数据
	if loginReq.Username == "admin" && loginReq.Password == "123456" {
		// 生成JWT token
		tokenString, err := middleware.GenerateToken(1, loginReq.Username)
		if err != nil {
			fmt.Println(999, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "Error creating token",
			})
			return
		}

		// 设置cookie
		c.SetCookie("token", tokenString, 24*60*60, "/", "", false, true)
		fmt.Println(333, tokenString)
		// 返回登录成功响应
		c.JSON(http.StatusOK, LoginResponse{
			Status:           "ok",
			Type:             loginReq.Type,
			CurrentAuthority: "admin",
		})
		return
	}
	fmt.Println(444)
	// 登录失败
	c.JSON(http.StatusOK, LoginResponse{
		Status:           "error",
		Type:             loginReq.Type,
		CurrentAuthority: "guest",
	})
}

func GetCurrentUser(c *gin.Context) {
	token := c.Query("token")
	fmt.Println(111, token)
	// 从上下文中获取用户信息
	claims, ok := middleware.GetUserFromContext(c)
	fmt.Println(111, claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "Unauthorized",
		})
		return
	}

	// TODO: 从数据库获取用户信息
	// 这里使用示例数据
	user := entity.User{
		ID:        claims.UserID,
		Username:  claims.Username,
		Avatar:    "https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png",
		Email:     "antdesign@alipay.com",
		Phone:     "0752-268888888",
		Address:   "西湖区工专路 77 号",
		Group:     "蚂蚁金服－某某某事业群－某某平台部－某某技术部",
		Access:    "admin",
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	fmt.Println(222, user)
	c.JSON(http.StatusOK, UserInfoResponse{
		Data: user,
	})
}

func Logout(c *gin.Context) {
	// 清除cookie
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
	})
}

// GetRuleList 获取规则列表
func GetRuleList(c *gin.Context) {
	// 这里可以从数据库获取数据，这里使用模拟数据
	rules := []entity.Rule{
		{
			Key:       99,
			Disabled:  false,
			Href:      "https://ant.design",
			Avatar:    "https://gw.alipayobjects.com/zos/rmsportal/udxAbMEhpwthVVcjLXik.png",
			Name:      "TradeCode 99",
			Owner:     "曲丽丽",
			Desc:      "这是一段描述",
			CallNo:    503,
			Status:    "0",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
			Progress:  81,
		},
		// ... 可以添加更多数据
	}

	response := RuleList{
		Data:     rules,
		Total:    len(rules),
		Success:  true,
		PageSize: 20,
		Current:  1,
	}

	c.JSON(http.StatusOK, response)
}
