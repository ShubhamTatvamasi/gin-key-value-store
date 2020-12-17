package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setup Router for Gin Server
func setupRouter() *gin.Engine {

	// Store all key value pairs in map
	store := map[string]string{}
	// List of all the subscribed keys
	subscription := make(map[string]struct{})

	// Route the gin web server
	r := gin.Default()

	// Serve static files
	r.Static("/assets", "./templates/assets")

	// Render index.html file
	r.LoadHTMLGlob("templates/index.html")

	// Home page
	r.GET("/", func(c *gin.Context) {

		// Sending empty string, so redering nothing.
		c.HTML(http.StatusOK, "index.html", gin.H{
			"": "",
		})

	})

	// Get function for getting a specfic value for a key
	r.GET("/get", func(c *gin.Context) {

		// Return Value
		c.String(200, store[c.Query("key")])

	})

	// Post function for adding or updating values
	r.POST("/post", func(c *gin.Context) {

		key := c.PostForm("key")
		value := c.PostForm("value")
		store[key] = value

		// Return Key=Value
		c.String(200, key+"="+value)

	})

	// Subscribe to keys for regular updates
	r.POST("/subscribe", func(c *gin.Context) {

		key := c.PostForm("key")
		subscription[key] = struct{}{}

		// Return subscribed key
		c.String(200, key)

	})

	// Keep connection alive for any updated on subscription
	r.GET("/stream", func(c *gin.Context) {

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

	})

	return r
}

func main() {
	// Start Gine Web Server
	r := setupRouter()
	r.Run(":80")
}
