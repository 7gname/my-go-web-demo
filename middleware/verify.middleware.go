package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err != nil || token == "" {
			c.Redirect(http.StatusMovedPermanently, "/403")
		}
		c.Next()
	}
}
