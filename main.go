package main

import (
	"yigouwang_api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// common.InitDB()
	r := gin.Default()
	r = routers.CollRouter(r)
	r.Run(":9999")
}
