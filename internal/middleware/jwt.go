package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your-secret-key") // 在实际应用中应该从配置文件读取

type Claims struct {
	UserID   uint   `json:"userid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// JWTAuth Gin中间件函数
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取token
		tokenString := extractToken(c)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Unauthorized",
			})
			c.Abort()
			return
		}

		// 解析token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Invalid token",
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user", claims)
		c.Next()
	}
}

// 从请求中提取token
func extractToken(c *gin.Context) string {
	// 1. 首先尝试从URL参数获取
	token := c.Query("token")
	if token != "" {
		return token
	}

	// 2. 然后尝试从Authorization头获取
	bearerToken := c.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	// 3. 最后尝试从cookie获取
	token, err := c.Cookie("token")
	if err == nil {
		return token
	}

	return ""
}

// GenerateToken 生成JWT token
func GenerateToken(userID uint, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// GetUserFromContext 从上下文中获取用户信息
func GetUserFromContext(c *gin.Context) (*Claims, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}
	claims, ok := user.(*Claims)
	return claims, ok
}
