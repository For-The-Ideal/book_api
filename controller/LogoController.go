package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuList struct {
	MenuList []MenuItem `json:"menuList"`
}

type MenuItem struct {
	Title    string    `json:"title"`
	Icon     string    `json:"icon"`
	FatherId int       `json:"fatherId"`
	SonList  []SonList `json:"sonList"`
}

type SonList struct {
	SonTitle  string `json:"sonTitle"`
	SonIcon   string `json:"sonIcon"`
	FatherId  int    `json:"fatherId"`
	SonId     int    `json:"sonId"`
	SonRouter string `json:"sonRouter"`
}

func Login(ctx *gin.Context) {
	var menuListInfo MenuList
	var sonList []SonList
	menuListInfo.MenuList = append(menuListInfo.MenuList, MenuItem{Title: "资源总揽", Icon: "md-desktop", FatherId: 1, SonList: sonList})
	sonList = append(sonList, SonList{SonTitle: "首页", SonId: 0, SonIcon: "ios-paper", FatherId: 1, SonRouter: "/home"})

	menuListInfo.MenuList = append(menuListInfo.MenuList, MenuItem{Title: "业务系统管理", Icon: "ios-keypad", FatherId: 2, SonList: sonList})
	sonList = append(sonList, SonList{SonTitle: "业务清单", SonId: 0, SonIcon: "md-list-box", FatherId: 2, SonRouter: "/articleList"})

	menuListInfo.MenuList = append(menuListInfo.MenuList, MenuItem{Title: "云主机管理", Icon: "md-phone-landscape", FatherId: 3, SonList: sonList})
	sonList = append(sonList, SonList{SonTitle: "网络部", SonId: 0, SonIcon: "md-list-box", FatherId: 3, SonRouter: "/videoList"})
	sonList = append(sonList, SonList{SonTitle: "信技部", SonId: 1, SonIcon: "md-list-box", FatherId: 3, SonRouter: ""})
	sonList = append(sonList, SonList{SonTitle: "采购部", SonId: 2, SonIcon: "md-list-box", FatherId: 3, SonRouter: ""})
	sonList = append(sonList, SonList{SonTitle: "办公室", SonId: 3, SonIcon: "md-list-box", FatherId: 3, SonRouter: ""})
	sonList = append(sonList, SonList{SonTitle: "其它部门", SonId: 4, SonIcon: "md-list-box", FatherId: 3, SonRouter: ""})

	// fmt.Println(menuListInfo, "menuListInfo")
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功",
		"code": http.StatusOK,
		"data": menuListInfo,
	})
}
