// main.go
package main

import (
	"benzinga/webhook/controllers"
	"benzinga/webhook/middleware"
	"benzinga/webhook/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    ginMode := os.Getenv("GIN_MODE") // Fetch GIN_MODE from env
    gin.SetMode(ginMode) 
    router := gin.Default()
    // Initialize services
    config := services.InitService(5) // Initialize the service with a default batch size of 5
    router.Use(func(c *gin.Context) {
        // Set config in Gin context
        c.Set("config", config)
        c.Next()
    })
    // Set up Gin and logging
   
    router.Use(middleware.RequestLogger())

    // Define routes
    router.GET("/healthz", controllers.HealthzHandler)
    router.POST("/log", controllers.LogHandler)

    // Start the server
    routerErr := router.Run(":8080")
	if routerErr != nil {
        log.Fatal("Error starting the server")
    }

}
