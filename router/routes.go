package router

import (
	"github.com/gin-gonic/gin"
	"occupie/controller/user"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/v1/auth", user.SignIn)

	return router
}
