// @file hourly_entity.go
// @desc hourly_entity
// @author KeeneChen <keenechen@qq.com>
// @since  2024.06.25-19:50:08

package entity

type CaiyunHourlyEntity struct {
	Status        string `json:"status"`      // "ok" 小时级预报状态
	Description   string `json:"description"` // 预报描述
	Precipitation []struct {
		Datatime    string  `json:"datatime"`
		Value       float64 `json:"value"`       // 降水量
		Probability float64 `json:"probability"` // 降水概率
	} `json:"precipitation"`
	Temperature []struct {
		Datatime string  `json:"datatime"`
		Value    float64 `json:"value"` // 温度
	} `json:"temperature"`
	ApparentTemperature []struct {
		Datatime string  `json:"datatime"`
		Value    float64 `json:"value"` // 体感温度
	} `json:"apparent_temperature"`
	Wind []struct {
		Datatime  string  `json:"datatime"`
		Speed     float64 `json:"speed"`     // 风速
		Direction float64 `json:"direction"` // 风向
	} `json:"wind"`
	Humidity []struct {
		Datatime string  `json:"datatime"`
		Value    float64 `json:"value"` // 相对湿度
	} `json:"humidity"`
	Cloudrate []struct {
		Datatime string  `json:"datatime"`
		Value    float64 `json:"value"` // 云量
	} `json:"cloudrate"`
	Skycon []struct {
		Datatime string `json:"datatime"`
		Value    string `json:"value"` // 天气现象
	} `json:"skycon"`
	Pressure []struct {
		Datatime string  `json:"datatime"`
		Value    float64 `json:"value"` // 气压
	} `json:"pressure"`
	Visibility []struct {
		Datatime string  `json:"datatime"`
		Value    float64 `json:"value"` // 能见度
	} `json:"visibility"`
	Dswrf []struct {
		Datatime string  `json:"datatime"`
		Value    float64 `json:"value"` // 短波辐射
	} `json:"dswrf"`
	AirQuality []struct {
		Aqi []struct {
			Datatime string `json:"datatime"`
			Value    struct {
				Chn float64 `json:"chn"` // 国内 AQI
				Usa int     `json:"usa"` // 美国 AQI
			} `json:"value"` // 空气质量指数
		} `json:"aqi"`
		Pm25 []struct {
			Datatime string  `json:"datatime"`
			Value    float64 `json:"value"` // PM2.5
		} `json:"pm25"`
	} `json:"air_quality"` // 空气质量
}
