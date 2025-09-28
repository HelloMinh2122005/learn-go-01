package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "minh.com/go-rest-gin-3/docs" // This is important for swagger docs
	"minh.com/go-rest-gin-3/internal/configs"
	"minh.com/go-rest-gin-3/internal/handlers/http"
	repositories "minh.com/go-rest-gin-3/internal/repositories/implementations"
	"minh.com/go-rest-gin-3/internal/routes"
	services "minh.com/go-rest-gin-3/internal/services/implementations"
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

	// 1. tạo ctx và cancel function: -> đảm bảo hủy context, connection đóng khi main kết thúc
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Bắt signal để graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
		cancel()
	}()

	// Khởi tạo MongoDB
	mongoDbConfig := configs.MongoConfig{
		URI:        "mongodb://localhost:27017",
		DBName:     "MongoDB",
		TimeoutSec: 10,
	}

	mongoClient, err := configs.NewMongo(ctx, mongoDbConfig)
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}

	defer func() {
		if err := mongoClient.Close(ctx); err != nil {
			log.Printf("error closing mongo: %v", err)
		}
	}()

	userRepo := repositories.NewUserRepository(mongoClient.DB.Collection("users"))
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
	srvErr := make(chan error)
	go func() {
		if err := router.Run(":8089"); err != nil {
			srvErr <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("Shutting down server.")
	case err := <-srvErr:
		log.Fatalf("server error: %v", err)
	}
}
