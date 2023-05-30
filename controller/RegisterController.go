package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {

	// fmt.Println(menuListInfo, "menuListInfo")
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功",
		"code": http.StatusOK,
		"data": "hellwo",
	})
}
