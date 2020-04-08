package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Err403(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Err: Err{
			Code: http.StatusForbidden,
			Msg:  http.StatusText(http.StatusForbidden),
		},
		Result: "",
	})
}

func Err404(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Err: Err{
			Code: http.StatusNotFound,
			Msg:  http.StatusText(http.StatusNotFound),
		},
		Result: "",
	})
}
