package services

import "time"

func InitService() {
    batchSize = getEnvAsInt("BATCH_SIZE", 5)
    postEndpoint = getEnv("POST_ENDPOINT", "https://eodlt9jg5ag9bze.m.pipedream.net")

    interval := getEnv("BATCH_INTERVAL", "10s")
    batchInterval, _ := time.ParseDuration(interval)

    go func() {
        ticker := time.NewTicker(batchInterval)
        for range ticker.C {
            sendBatch()
        }
    }()
}