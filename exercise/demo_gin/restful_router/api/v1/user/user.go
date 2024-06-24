package user

import "github.com/gin-gonic/gin"

func HandlerPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}

func HandlerLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func HandlerRegister(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register",
	})
}
