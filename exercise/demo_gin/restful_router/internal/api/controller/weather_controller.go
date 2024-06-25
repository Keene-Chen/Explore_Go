package controller

import (
	"Explore_Go/exercise/demo_gin/restful_router/internal/api/request"
	"Explore_Go/exercise/demo_gin/restful_router/internal/api/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// HandlerWeather 处理天气请求
func HandlerWeather(c *gin.Context) {
	wr := request.WeatherRequest{
		BaseURL: "https://api.caiyunapp.com/v2.6",
	}
	if err := c.ShouldBindUri(&wr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	fmt.Printf("Request URL: %s\n\n", wr.ToString())

	crr := &response.CaiyunRealtimeResponse{}
	client := resty.New()
	resp, err := client.R().Get(wr.ToString())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to get weather data"})
		return
	}
	fmt.Printf("Response Body: %+v\n", crr)

	c.Data(http.StatusOK, "application/json; charset=utf-8", resp.Body())
}
