package router

import (
	apiUser "Explore_Go/exercise/demo_gin/restful_router/api/v1/user"
	apiWeather "Explore_Go/exercise/demo_gin/restful_router/api/v1/weather"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", apiUser.HandlerPing)

	v1 := r.Group("api/v1/")
	{
		user := v1.Group("user/")
		{
			user.POST("/register", apiUser.HandlerRegister)
			user.GET("/login", apiUser.HandlerLogin)
		}

		weather := v1.Group("weather/")
		{
			weather.GET("/:token/:lnglat/:mode", apiWeather.HandlerWeather)
		}
	}

	return r
}
