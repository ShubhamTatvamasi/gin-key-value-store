package router

import (
	"github.com/gin-gonic/gin"
)

func subscribe(c *gin.Context) {

	key := c.PostForm("key")
	subscription[key] = struct{}{}

	// Return subscribed key
	c.String(200, key)
}
