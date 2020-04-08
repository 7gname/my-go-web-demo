package main

import (
	"my-go-web-demo/ctrl"
	"my-go-web-demo/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	//中间件函数
	setMiddleWare(e)
	//注册路由
	setHandler(e)

	e.Run(":8080")
}

func setMiddleWare(e *gin.Engine) {
	//捕获404跳转到404公益页
	e.Use(middleware.Err404())
}

func setHandler(e *gin.Engine) {
	e.GET("/403", ctrl.Err403)
	e.GET("/404", ctrl.Err404)

	user := e.Group("/user")
	//Parameters in path
	user.GET("/info/:name", ctrl.GetUserInfo)
	user.POST("/add", ctrl.AddUser)

	comment := e.Group("/comment", middleware.Verify())
	comment.POST("/add", ctrl.AddComment)
}
