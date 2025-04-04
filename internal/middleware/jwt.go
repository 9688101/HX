package middleware

import (
	"net/http"
	"strings"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/pkg/blacklist"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// CustomClaims 自定义 JWT Claims
type CustomClaims struct {
	Username string `json:"username"`
	Role     int    `json:"role"`
	ID       int    `json:"id"`
	Status   int    `json:"status"`
	jwt.RegisteredClaims
}

// jwtAuthHelper 解析 JWT 并判断权限，minRole 为最低权限要求
func jwtAuthHelper(c *gin.Context, minRole int) {
	// 从 Authorization 头中获取 token，通常格式为 "Bearer token..."
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Missing Authorization header",
		})
		c.Abort()
		return
	}

	// 解析 Bearer token
	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Token not provided",
		})
		c.Abort()
		return
	}

	// 解析 token，使用配置中的 JWT Secret
	secret := config.GetServerConfig().JWTSecret // 假设该函数返回 JWT secret 字符串
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid token",
		})
		c.Abort()
		return
	}

	// 转换 claims 为自定义类型
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid token claims",
		})
		c.Abort()
		return
	}

	// 判断用户状态是否正常，以及是否被封禁
	if claims.Status == entity.UserStatusDisabled || blacklist.IsUserBanned(claims.ID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户已被封禁",
		})
		c.Abort()
		return
	}

	// 判断用户权限是否满足要求
	if claims.Role < minRole {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无权进行此操作，权限不足",
		})
		c.Abort()
		return
	}

	// 将用户信息存入 Gin 的 Context，方便后续使用
	c.Set("username", claims.Username)
	c.Set("role", claims.Role)
	c.Set("id", claims.ID)
	c.Next()
}

// JWTUserAuth 针对普通用户的 JWT 认证中间件
func JWTUserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtAuthHelper(c, entity.RoleCommonUser)
	}
}

// JWTAdminAuth 针对管理员的 JWT 认证中间件
func JWTAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtAuthHelper(c, entity.RoleAdminUser)
	}
}

// JWTRootAuth 针对超级管理员的 JWT 认证中间件
func JWTRootAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtAuthHelper(c, entity.RoleRootUser)
	}
}
