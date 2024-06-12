package router

import (
	_ "restful_api/app/controller"

	ginAutoRouter "github.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//初始化路由
	r := gin.Default()

	//绑定基本路由，访问路径：/User/List
	ginAutoRouter.Bind(r)

	return r
}
