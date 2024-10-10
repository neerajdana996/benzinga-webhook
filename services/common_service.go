package services

import (
	"benzinga/webhook/models"
	"log"
	"time"
)

type ServiceConfig struct {
    BatchSize    int
    PostEndpoint string
    BatchInterval time.Duration
    Cache        []models.Payload
}

// InitService initializes the service configuration and starts the batch processing routine
func InitService(customBatchSize int) *ServiceConfig {
    batchSize := getEnvAsInt("BATCH_SIZE", customBatchSize) // Default batch size: 5
    postEndpoint := getEnv("POST_ENDPOINT", "https://eodlt9jg5ag9bze.m.pipedream.net") // Default endpoint

    interval := getEnv("BATCH_INTERVAL", "10s") // Default batch interval: 10s
    batchInterval, err := time.ParseDuration(interval)
    if err != nil {
        log.Fatalf("Failed to parse batch interval: %v", err) // Terminate if parsing fails
    }

    config := &ServiceConfig{
        BatchSize:    batchSize,
        PostEndpoint: postEndpoint,
        BatchInterval: batchInterval,
    }

    // Start the ticker to trigger batch sending at defined intervals
    go func() {
        ticker := time.NewTicker(config.BatchInterval)
        defer ticker.Stop()
        
        for range ticker.C {
            sendBatch(config) // Pass config to sendBatch function
        }
    }()

    log.Printf("Service initialized with batch size: %d, post endpoint: %s, batch interval: %s", 
        config.BatchSize, config.PostEndpoint, config.BatchInterval)

    return config
}