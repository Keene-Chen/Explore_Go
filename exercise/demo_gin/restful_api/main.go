// @file main.go
// @desc gin main
// @author KeeneChen <keenechen@qq.com>
// @since  2024.06.04-10:52:23

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/get/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"username": "zhangsan",
		})
	})

	r.POST("/post/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"username": "lisi",
		})
	})

	r.PUT("/put/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"username": "wangwu",
		})
	})

	r.DELETE("/delete/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"username": "zhaoliu",
		})
	})

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}
