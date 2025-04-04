package middleware

import (
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/pkg/blacklist"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"net/http"
)

func authHelper(c *gin.Context, minRole int) {
	session := sessions.Default(c)
	username := session.Get("username")
	role := session.Get("role")
	id := session.Get("id")
	status := session.Get("status")
	if username == nil {
		// Check access token
		accessToken := c.Request.Header.Get("Authorization")
		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "无权进行此操作，未登录且未提供 access token",
			})
			c.Abort()
			return
		}
	}
	if status.(int) == entity.UserStatusDisabled || blacklist.IsUserBanned(id.(int)) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户已被封禁",
		})
		session := sessions.Default(c)
		session.Clear()
		_ = session.Save()
		c.Abort()
		return
	}
	if role.(int) < minRole {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无权进行此操作，权限不足",
		})
		c.Abort()
		return
	}
	c.Set("username", username)
	c.Set("role", role)
	c.Set("id", id)
	c.Next()
}

func UserAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHelper(c, entity.RoleCommonUser)
	}
}

func AdminAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHelper(c, entity.RoleAdminUser)
	}
}

func RootAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHelper(c, entity.RoleRootUser)
	}
}

// Auth 中间件：基于 Session 和 JWT 的混合认证示例
// 此处为示例，如果同时支持 session 和 JWT，可先检查 Session，再检查 JWT
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		role := session.Get("role")
		id := session.Get("id")
		status := session.Get("status")
		// 如果 session 中未找到用户信息，则尝试检查 JWT（可以扩展此处）
		if username == nil {
			accessToken := c.Request.Header.Get("Authorization")
			if accessToken == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"success": false,
					"message": "未登录且未提供 access token",
				})
				c.Abort()
				return
			}
			// 此处建议调用 JWT 中间件来解析 token，并设置对应的用户信息
			// 例如：jwtAuthHelper(c, entity.RoleCommonUser)
		}

		// 检查用户状态与权限（假设 status 为 int）
		if statusInt, ok := status.(int); ok && (statusInt == entity.UserStatusDisabled || blacklist.IsUserBanned(id.(int))) {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "用户已被封禁",
			})
			session.Clear()
			_ = session.Save()
			c.Abort()
			return
		}
		// 此处可进一步判断 role 等权限信息
		c.Set("username", username)
		c.Set("role", role)
		c.Set("id", id)
		c.Next()
	}
}
