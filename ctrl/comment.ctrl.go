package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddComment(c *gin.Context) {
	commentator := c.PostForm("commentator")
	content := c.PostForm("content")
	c.JSON(http.StatusOK, Response{
		Err: NewErr(SuccessCode, SuccessCodeMsg),
		Result: struct {
			Commentator string
			Content     string
		}{
			commentator,
			content,
		},
	})
}
