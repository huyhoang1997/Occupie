package router

import (
	"github.com/gin-gonic/gin"
	"occupie/controller"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/v1/auth", controller.SignIn)

	return router
}
