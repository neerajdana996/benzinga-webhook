package services

import (
	"benzinga/webhook/models"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAddToCache(t *testing.T) {
    // Initialize config using initConfig
    config := InitService(0)
    config.BatchSize = 3 // Override batch size for the test

    payload := models.Payload{
        UserID:    1,
        Total:     10.5,
        Title:     "Test payload",
        Meta:      models.Meta{},
        Completed: false,
    }

    // Add three payloads and check if batch is sent
    AddToCache(payload, config)
    AddToCache(payload, config)
    AddToCache(payload, config)

    // Test that the cache has been cleared after batch processing
    assert.Equal(t, 0, len(config.Cache))
}

func TestSendBatch(t *testing.T) {
    // Initialize config using initConfig
    config := InitService(0)
    config.BatchSize = 2 // Override batch size for the test

    payload := models.Payload{
        UserID:    1,
        Total:     10.5,
        Title:     "Test payload",
        Meta:      models.Meta{},
        Completed: false,
    }

    // Add two payloads to trigger batch sending
    AddToCache(payload, config)
    AddToCache(payload, config)

    // Cache should be cleared after sending the batch
    assert.Equal(t, 0, len(config.Cache))
}

func TestSendBatch_Mock(t *testing.T) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    // Mock the POST request
    httpmock.RegisterResponder("POST", "https://eodlt9jg5ag9bze.m.pipedream.net",
        httpmock.NewStringResponder(200, `OK`))

    // Initialize config using initConfig
    config := InitService(0)

    payload := models.Payload{
        UserID:    1,
        Total:     10.5,
        Title:     "Mocked payload",
        Meta:      models.Meta{},
        Completed: false,
    }

    // Add payloads to the cache
    config.Cache = append(config.Cache, payload, payload)

    // Send the batch
    sendBatch(config)

    // Verify that the cache was cleared and the mock was called
    assert.Equal(t, 0, len(config.Cache))
    info := httpmock.GetCallCountInfo()
    assert.Equal(t, 1, info["POST https://eodlt9jg5ag9bze.m.pipedream.net"])
}
