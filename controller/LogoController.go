package controller

import (
	"chatGpt_api/common"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义user类型结构体
type User struct {
	id         int    `json:"id"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	Head_photo string `json:"head_photo"`
}

func Login(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")

	fmt.Println(account, "account")
	fmt.Println(password, "password")
	userInfo, err := QueryUserByCode(account, password)
	if !err {
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusAccepted,
			"data":   gin.H{},
			"msg":    "账号密码不存在",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   userInfo,
		"msg":    "success",
	})

}

func QueryUserByCode(account string, password string) (*User, bool) {
	u := &User{}             // 定义User切片
	db := common.ConnMySQL() // 调用db包ConnMySQL()
	// 预编译查询sql创建 statement
	stmt, err := db.Prepare("SELECT account, password FROM `user` WHERE account = ? && password = ? && id > 0")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer stmt.Close()
	// 执行查询sql，返回查询结果rows
	rows, err := stmt.Query(account, password)
	if err != nil {
		// 打印错误信息
		log.Fatal(err)
		// 抛出错误信息，阻止程序继续运行
		return u, false
	}
	fmt.Print(rows, "rows")
	// 遍历rows
	for rows.Next() {
		fmt.Println(u, "u")
		// 扫描rows的每一列并保存数据到User对应字段
		err := rows.Scan(&u.Account, &u.Password)
		// &u.id,
		// , &u.Password
		if err != nil {
			// 打印错误信息
			log.Fatal(err)
			// 抛出错误信息，阻止程序继续运行
			panic(err)
		}
		// 扫描后的user加入到切片
		// s = append(s, u)
	}
	return u, true
}
