package router

import (
	"Explore_Go/exercise/demo_gin/restful_router/internal/api/controller"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controller.HandlerPing)

	v1 := r.Group("api/v1/")
	{
		user := v1.Group("user/")
		{
			user.POST("/register", controller.HandlerRegister)
			user.GET("/login", controller.HandlerLogin)
		}

		weather := v1.Group("weather/")
		{
			weather.GET("/:token/:lnglat/:mode", controller.HandlerWeather)
		}
	}

	return r
}
