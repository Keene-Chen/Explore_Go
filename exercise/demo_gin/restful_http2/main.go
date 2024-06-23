package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.RunTLS("127.0.0.1:443", "./key/server.crt", "./key/server.key"); err != nil {
		panic(err)
	}
}
