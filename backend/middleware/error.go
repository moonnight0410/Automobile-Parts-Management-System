package middleware

import (
	"net/http"

	"automobile-parts-backend/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 {
			return
		}
		err := c.Errors.Last().Err
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
	}
}
