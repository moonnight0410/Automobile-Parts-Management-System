package middleware

import (
	"net/http"
	"strings"

	"automobile-parts-backend/config"
	"automobile-parts-backend/utils"
	"github.com/gin-gonic/gin"
)

// Auth JWT认证中间件
// 验证请求头中的JWT令牌，并将用户信息设置到上下文
// 参数：
//   - cfg: 应用配置，包含JWT密钥
// 返回：
//   - gin.HandlerFunc: Gin中间件函数
func Auth(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从请求头获取Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.Error(http.StatusUnauthorized, "missing authorization header"))
			c.Abort() // 终止请求处理
			return
		}
		
		// 2. 验证Authorization头格式（必须是Bearer token）
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, utils.Error(http.StatusUnauthorized, "invalid authorization header"))
			c.Abort()
			return
		}
		
		// 3. 解析并验证JWT令牌
		claims, err := utils.ParseToken(parts[1], cfg.JWTSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.Error(http.StatusUnauthorized, "invalid token"))
			c.Abort()
			return
		}
		
		// 4. 将用户信息设置到上下文，供后续处理函数使用
		c.Set("userID", claims.UserID) // 用户ID
		c.Set("role", claims.Role)     // 用户角色（manufacturer/automaker/aftersale）
		
		// 5. 继续执行后续中间件和处理函数
		c.Next()
	}
}
