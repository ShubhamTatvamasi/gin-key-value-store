package router

import (
	"embed"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var templates embed.FS

// Store all key value pairs in map
var store = map[string]string{}

// List of all the subscribed keys
var subscription = make(map[string]struct{})

// Setup Router for Gin Server
func SetupRouter() *gin.Engine {

	// Route the gin web server
	r := gin.Default()

	// Serve static files
	r.Static("/assets", "./templates/assets")

	// Home page
	r.GET("/", root)

	// Get function for getting a specfic value for a key
	r.GET("/get", get)

	// Post function for adding or updating values
	r.POST("/post", post)

	// Subscribe to keys for regular updates
	r.POST("/subscribe", subscribe)

	// Unsubscribe key
	r.POST("/unsubscribe", unsubscribe)

	// Keep connection alive for any updated on subscription
	r.GET("/stream", stream)

	return r
}
