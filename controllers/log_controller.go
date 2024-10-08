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
    var payload models.Payload
    if err := c.ShouldBindJSON(&payload); err != nil {
        logrus.WithError(err).Error("Invalid JSON received")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    services.AddToCache(payload) // Add the payload to the service cache
    logrus.WithFields(logrus.Fields{"user_id": payload.UserID, "title": payload.Title}).Info("Payload received")

    c.JSON(http.StatusOK, gin.H{"status": "logged"})
}
