package router

import (
	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {

	// Return Value
	c.String(200, store[c.Query("key")])
}
