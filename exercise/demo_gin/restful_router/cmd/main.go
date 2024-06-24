package main

import (
	"Explore_Go/exercise/demo_gin/restful_router/router"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	gin.SetMode(gin.DebugMode)
	r := router.RouterSetup()

	certPath, _ := filepath.Abs("key/server.crt")
	keyPath, _ := filepath.Abs("key/server.key")
	r.RunTLS("127.0.0.1:8080", certPath, keyPath)
}
