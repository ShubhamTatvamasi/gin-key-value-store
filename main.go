package main

import (
	"github.com/ShubhamTatvamasi/gin-key-value-store/routes"
)

func main() {
	// Start Gine Web Server
	r := routes.SetupRouter()
	r.Run(":80")
}
