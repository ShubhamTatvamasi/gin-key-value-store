package router

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func stream(c *gin.Context) {
	// Add header for setting Server Sent Events to live stream
	w := c.Writer
	w.Header().Set("Content-Type", "text/event-stream")

	keyValues := make(map[string]interface{})

	// Collect of subscribed key values
	for key := range subscription {
		keyValues[key] = store[key]
	}

	// Package all keys in JSON
	empData, _ := json.Marshal(keyValues)
	jsonString := string(empData)

	// Send JSON back to live stream
	w.Write([]byte("data:" + jsonString + "\n\n"))
}
