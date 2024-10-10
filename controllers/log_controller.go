// controllers/log_controller.go
package controllers

import (
	"benzinga/webhook/models"
	"benzinga/webhook/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HealthzHandler(c *gin.Context) {
    c.String(http.StatusOK, "OK")
}

func LogHandler(c *gin.Context) {
    configInterface, exists := c.Get("config")
    if !exists {
        logrus.Error("Config not found in context")
        c.JSON(500, gin.H{"error": "Internal server error"})
        return
    }

    // Type assert the config to its original type
    config, ok := configInterface.(*services.ServiceConfig)
    if !ok {
        logrus.Error("Failed to cast config from context")
        c.JSON(500, gin.H{"error": "Internal server error"})
        return
    }
    var payload models.Payload
    if err := c.ShouldBindJSON(&payload); err != nil {
        logrus.WithError(err).Error("Invalid JSON received")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    services.AddToCache(payload,config) // Add the payload to the service cache
    c.JSON(http.StatusOK, gin.H{"status": "logged"})
}
