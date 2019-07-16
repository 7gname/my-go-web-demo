package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Err404() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Writer.Status() == http.StatusNotFound {
			c.Redirect(http.StatusMovedPermanently, "/404")
		}
		c.Next()
	}
}
