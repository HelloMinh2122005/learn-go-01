package main

import (
	http "minh.com/go-rest-gin-3/internal/handlers/http"
	repositories "minh.com/go-rest-gin-3/internal/repositories/implementations"
	routes "minh.com/go-rest-gin-3/internal/routes"
	services "minh.com/go-rest-gin-3/internal/services/implementations"

	gin "github.com/gin-gonic/gin"
)

func main() {
	userRepo := repositories.NewUserRepository()
	// TODO: delete this shit
	userRepo.SeedDummyData()
	userService := services.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)

	// Khởi tạo Gin router
	router := gin.Default()

	// Đăng ký routes
	routes.RegisterUserRoutes(router, userHandler)

	// Chạy server
	router.Run(":8089")
}
