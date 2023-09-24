package main

import (
	"github.com/ShubhamTatvamasi/gin-key-value-store/router"
)

func main() {
	// Start Gin Web Server
	r := router.SetupRouter()
	r.Run(":80")
}
