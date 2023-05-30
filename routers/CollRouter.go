package routers

import (
	"chatGpt_api/controller"

	"github.com/gin-gonic/gin"
)

func CollRouter(r *gin.Engine) *gin.Engine {
	userApi := r.Group("/api")
	{
		userApi.GET("/auth/verificationCode", controller.VerificationCode)
		userApi.POST("/auth/login", controller.Login)
		userApi.POST("/auth/register", controller.Register)
	}
	return r
}
