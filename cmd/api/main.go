package main

import (
	http "minh.com/go-rest-gin-3/internal/handlers/http"
	repositories "minh.com/go-rest-gin-3/internal/repositories/implementations"
	routes "minh.com/go-rest-gin-3/internal/routes"
	services "minh.com/go-rest-gin-3/internal/services/implementations"

	gin "github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "minh.com/go-rest-gin-3/docs" // Import to embed Swagger docs
)

// @title test API
// @version 1.0
// @description This is a sample server for a test service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8089
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

	// Đăng ký route swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8089/swagger/doc.json"),
	))

	// Chạy server
	router.Run(":8089")
}
