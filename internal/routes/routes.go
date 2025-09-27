package routes

import (
	"minh.com/go-rest-gin-3/internal/handlers/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userHandler *http.UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("", userHandler.GetAllUsers)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.POST("", userHandler.CreateUser)
	}
}
