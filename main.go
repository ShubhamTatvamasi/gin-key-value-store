package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// Store all key value pairs in map
	store := map[string]string{}

	// Route the gin web server
	r := gin.Default()

	// Get function for getting a specfic value for a key
	r.GET("/", func(c *gin.Context) {

		c.String(200, store[c.Query("key")]+"\n")

	})

	// Post function for adding or updating values
	r.POST("/", func(c *gin.Context) {

		key := c.PostForm("key")
		value := c.PostForm("value")
		store[key] = value

		c.String(200, key+"="+value+"\n")

	})

	r.Run(":80")
}
