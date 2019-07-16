package ctrl

import (
	"demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	name := c.Param("name")
	userInfo, err := service.GetUserInfo(name)
	if err != nil {
		c.JSON(
			http.StatusOK,
			Response{
				Err:    NewErr(FailedCode, FailedCodeMsg),
				Result: "",
			},
		)
	}
	c.JSON(
		http.StatusOK,
		Response{
			Err:    NewErr(SuccessCode, SuccessCodeMsg),
			Result: userInfo,
		},
	)
}

func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	c.String(http.StatusOK, "add %s", name)
}
