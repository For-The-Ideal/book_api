package main

import (
	"chatGpt_api/common"
	"chatGpt_api/middleware"
	"chatGpt_api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	common.ConnMySQL()
	r := gin.Default()
	r.Use(middleware.Cors())
	r = routers.CollRouter(r)
	r.Run(":9999")
}
