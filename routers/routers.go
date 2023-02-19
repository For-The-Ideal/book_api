package routers

import (
	"yigouwang_api/controller"

	"github.com/gin-gonic/gin"
)

func CollRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/login", controller.Login)
	return r
}
