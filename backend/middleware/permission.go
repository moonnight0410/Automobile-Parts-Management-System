package middleware

import (
	"net/http"

	"automobile-parts-backend/utils"

	"github.com/gin-gonic/gin"
)

func Permission(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, utils.Error(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}
		role, ok := roleValue.(string)
		if !ok || role == "" {
			c.JSON(http.StatusUnauthorized, utils.Error(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}
		if role != requiredRole {
			c.JSON(http.StatusForbidden, utils.Error(http.StatusForbidden, "forbidden"))
			c.Abort()
			return
		}
		c.Next()
	}
}
