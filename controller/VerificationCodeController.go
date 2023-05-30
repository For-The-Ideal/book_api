package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerificationCode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "注册成功",
		"code": http.StatusOK,
		"data": [...]int{1, 2, 3, 4, 5, 6, 7, 8},
	})
}
