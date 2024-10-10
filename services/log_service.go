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

// ServiceConfig holds the configuration for batch processing

// AddToCache adds a payload to the in-memory cache
func AddToCache(payload models.Payload, config *ServiceConfig) {
    config.Cache = append(config.Cache, payload)

    if len(config.Cache) >= config.BatchSize {
        sendBatch(config)
    }
}

// sendBatch sends the batched payloads to the external endpoint
func sendBatch(config *ServiceConfig) {
    if len(config.Cache) == 0 {
        return
    }

    logrus.WithFields(logrus.Fields{
        "batch_size": len(config.Cache),
    }).Info("Sending batch...")

    // Serialize cache to JSON
    jsonData, err := json.Marshal(config.Cache)
    if err != nil {
        logrus.WithError(err).Error("Failed to serialize batch")
        return
    }

    // Retry logic
    success := false
    for i := 0; i < 3; i++ {
        resp, err := http.Post(config.PostEndpoint, "application/json", bytes.NewBuffer(jsonData))
        if err == nil && resp.StatusCode == http.StatusOK {
            logrus.WithFields(logrus.Fields{
                "batch_size": len(config.Cache),
                "status":     resp.StatusCode,
            }).Info("Batch successfully sent")
            success = true
            break
        }

        logrus.WithFields(logrus.Fields{
            "attempt": i + 1,
        }).Warnf("Failed to send batch, retrying... %v", err)
        time.Sleep(2 * time.Second)
    }

    if !success {
        logrus.Error("Failed to send batch after 3 attempts, exiting")
        os.Exit(1)
    }

    // Clear cache after successful send
    config.Cache = []models.Payload{}
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
