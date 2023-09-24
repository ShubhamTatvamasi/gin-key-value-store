package router

import (
	"github.com/gin-gonic/gin"
)

func post(c *gin.Context) {

	key := c.PostForm("key")
	value := c.PostForm("value")
	store[key] = value

	// Return Key=Value
	c.String(200, key+"="+value)
}
