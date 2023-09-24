package router

import (
	"github.com/gin-gonic/gin"
)

func unsubscribe(c *gin.Context) {

	key := c.PostForm("key")
	delete(subscription, key)

	// Return unsubscribed key
	c.String(200, key)
}
