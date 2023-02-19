package routers

import (
	"yigouwang_api/controller"

	"github.com/gin-gonic/gin"
)

func CollRouter(r *gin.Engine) *gin.Engine {
	r.GET("/api/auth/verificationCode", controller.VerificationCode)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/register", controller.Register)
	return r
}
