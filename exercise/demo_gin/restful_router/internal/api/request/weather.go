package request

import "fmt"

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
