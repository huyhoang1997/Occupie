package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(c *gin.Context) {
	//headers := request.PrepareHeaders(c.Request.Header)
	c.JSON(http.StatusOK, "ok")
}
