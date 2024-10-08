// main.go
package main

import (
	"benzinga/webhook/controllers"
	"benzinga/webhook/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	gin.SetMode(gin.ReleaseMode)
    // Initialize services
    services.InitService()

    // Set up Gin and logging
    router := gin.Default()
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.Info("Starting application...")

    // Define routes
    router.GET("/healthz", controllers.HealthzHandler)
    router.POST("/log", controllers.LogHandler)

    // Start the server
    routerErr := router.Run(":8080")
	if routerErr != nil {
        log.Fatal("Error starting the server")
    }

}
