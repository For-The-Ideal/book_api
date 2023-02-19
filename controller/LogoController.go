package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功",
		"code": http.StatusOK,
	})
}
