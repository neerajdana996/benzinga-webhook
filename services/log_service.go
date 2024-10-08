// services/log_service.go
package services

import (
	"benzinga/webhook/models"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var cache []models.Payload
var batchSize int
var postEndpoint string
var GIN_MODE string

func InitServiceForTest(customBatchSize int) {
    batchSize = customBatchSize
    cache = []models.Payload{}
    postEndpoint = "https://eodlt9jg5ag9bze.m.pipedream.net" // Mock endpoint for testing,
	GIN_MODE = "release"
}

// AddToCache adds a payload to the in-memory cache
func AddToCache(payload models.Payload) {
    cache = append(cache, payload)

    if len(cache) >= batchSize {
        sendBatch()
    }
}

// sendBatch sends the batched payloads to the external endpoint
func sendBatch() {
    if len(cache) == 0 {
        return
    }

    logrus.WithFields(logrus.Fields{
        "batch_size": len(cache),
    }).Info("Sending batch...")

    // Serialize cache to JSON
    jsonData, err := json.Marshal(cache)
    if err != nil {
        logrus.WithError(err).Error("Failed to serialize batch")
        return
    }

    // Retry logic
    success := false
    for i := 0; i < 3; i++ {
        resp, err := http.Post(postEndpoint, "application/json", bytes.NewBuffer(jsonData))
        if err == nil && resp.StatusCode == http.StatusOK {
            logrus.WithFields(logrus.Fields{
                "batch_size": len(cache),
                "status":     resp.StatusCode,
            }).Info("Batch successfully sent")
            success = true
            break
        }

        logrus.WithFields(logrus.Fields{
            "attempt": i + 1,
        }).Warn("Failed to send batch, retrying...",err)
        time.Sleep(2 * time.Second)
    }

    if !success {
        logrus.Error("Failed to send batch after 3 attempts, exiting")
        os.Exit(1)
    }

    // Clear cache after successful send
    cache = []models.Payload{}
}

// Helper functions to get environment variables with defaults
func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func getEnvAsInt(key string, defaultValue int) int {
    valueStr := os.Getenv(key)
    value, err := strconv.Atoi(valueStr)
    if err != nil {
        return defaultValue
    }
    return value
}
