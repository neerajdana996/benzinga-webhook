package controllers

import (
	"benzinga/webhook/models"
	"benzinga/webhook/services"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Helper function to initialize the Gin router with middleware
func setupRouter(config *services.ServiceConfig) *gin.Engine {
    router := gin.Default()
    
    // Middleware to inject config into the Gin context
    router.Use(func(c *gin.Context) {
        c.Set("config", config)
        c.Next()
    })

    // Define routes
    router.GET("/healthz", HealthzHandler)
    router.POST("/log", LogHandler)
    
    return router
}

func TestHealthzHandler(t *testing.T) {
    // Set up the Gin context and recorder
    router := setupRouter(services.InitService(5)) // Use the InitService function to create config

    req, _ := http.NewRequest("GET", "/healthz", nil)
    w := httptest.NewRecorder()

    // Perform the request
    router.ServeHTTP(w, req)

    // Assertions
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, "OK", w.Body.String())
}

func TestLogHandler_ValidPayload(t *testing.T) {
    config := services.InitService(3) // Initialize the service with a batch size of 3
    router := setupRouter(config)

    payload := models.Payload{
        UserID:    1,
        Total:     1.65,
        Title:     "Test",
        Meta:      models.Meta{},
        Completed: false,
    }

    payloadJSON, _ := json.Marshal(payload)
    req, _ := http.NewRequest("POST", "/log", bytes.NewBuffer(payloadJSON))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "logged")
}

func TestLogHandler_InvalidPayload(t *testing.T) {
    config := services.InitService(3) // Initialize the service with a batch size of 3
    router := setupRouter(config)

    invalidJSON := `{"user_id": "invalid"}` // invalid JSON (user_id should be int)
    req, _ := http.NewRequest("POST", "/log", bytes.NewBuffer([]byte(invalidJSON)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Contains(t, w.Body.String(), "Invalid JSON")
}
