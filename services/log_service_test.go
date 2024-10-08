package services

import (
	"benzinga/webhook/models"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAddToCache(t *testing.T) {
    InitServiceForTest(3) // Initialize the service with a batch size of 3

    payload := models.Payload{
        UserID:    1,
        Total:     10.5,
        Title:     "Test payload",
        Meta:      models.Meta{},
        Completed: false,
    }

    // Add three payloads and check if batch is sent
    AddToCache(payload)
    AddToCache(payload)
    AddToCache(payload)

	println(len(cache))
    // Test that the cache has been cleared after batch processing
    assert.Equal(t, 0, len(cache))
}

func TestSendBatch(t *testing.T) {
    InitServiceForTest(2) // Initialize with batch size of 2

    payload := models.Payload{
        UserID:    1,
        Total:     10.5,
        Title:     "Test payload",
        Meta:      models.Meta{},
        Completed: false,
    }

    // Add two payloads to trigger batch sending
    AddToCache(payload)
    AddToCache(payload)

    assert.Equal(t, 0, len(cache)) // Cache should be cleared after sending the batch
}

func TestSendBatch_Mock(t *testing.T) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    // Mock the POST request
    httpmock.RegisterResponder("POST", "https://eodlt9jg5ag9bze.m.pipedream.net",
        httpmock.NewStringResponder(200, `OK`))

    payload := models.Payload{
        UserID:    1,
        Total:     10.5,
        Title:     "Mocked payload",
        Meta:      models.Meta{},
        Completed: false,
    }

    cache = append(cache, payload, payload)

    sendBatch()

    // Verify that the cache was cleared and the mock was called
    assert.Equal(t, 0, len(cache))
    info := httpmock.GetCallCountInfo()
    assert.Equal(t, 1, info["POST https://eodlt9jg5ag9bze.m.pipedream.net"])
}