package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func root(c *gin.Context) {

	// Read the contents of index.html from the embedded filesystem
	indexHTML, err := templates.ReadFile("templates/index.html")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Sending empty string, so rendering nothing.
	c.Data(http.StatusOK, "text/html; charset=utf-8", indexHTML)
}
