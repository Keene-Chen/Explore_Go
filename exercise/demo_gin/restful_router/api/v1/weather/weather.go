package weather

import (
	entity "Explore_Go/exercise/demo_gin/restful_router/api/entity/weather"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// WeatherRequest 定义请求参数结构体
type WeatherRequest struct {
	BaseURL string
	Token   string `uri:"token" binding:"required"`
	Lnglat  string `uri:"lnglat" binding:"required"`
	Mode    string `uri:"mode" binding:"required"`
}

// ToString 方法生成请求的完整URL
func (wr *WeatherRequest) ToString() string {
	return fmt.Sprintf("%s/%s/%s/%s", wr.BaseURL, wr.Token, wr.Lnglat, wr.Mode)
}

// HandlerWeather 处理天气请求
func HandlerWeather(c *gin.Context) {
	wr := WeatherRequest{
		BaseURL: "https://api.caiyunapp.com/v2.6",
	}
	if err := c.ShouldBindUri(&wr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	fmt.Printf("Request URL: %s\n\n", wr.ToString())

	caiyunResponse := &entity.CaiyunResponse{}
	client := resty.New()
	resp, err := client.R().SetResult(caiyunResponse).Get(wr.ToString())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to get weather data"})
		return
	}
	fmt.Printf("Response Body: %+v\n", caiyunResponse)

	c.Data(http.StatusOK, "application/json; charset=utf-8", resp.Body())
}
