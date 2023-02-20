package controller

import (
	"fmt"
	"net/http"
	"yigouwang_api/common"
	"yigouwang_api/util"

	"github.com/gin-gonic/gin"
)

type User struct {
	Account    string `json:"account"`
	Password   string `json:"password"`
	CreateTime string `json:create_time`

	AccountType  int    `json:account_type`
	HeadPhoto    string `json:head_photo`
	RegisterType int    `json:register_type`
}

func Register(ctx *gin.Context) {
	db := common.GetDB()
	user := &User{
		Account:      ctx.PostForm("account"),
		Password:     util.HashPassword(ctx.PostForm("password")),
		CreateTime:   util.GetTime(),
		AccountType:  1,
		HeadPhoto:    "",
		RegisterType: 1,
	}
	err := db.Create(user).Error
	if err != nil {
		fmt.Println("insert Error")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": "",
	})
}
