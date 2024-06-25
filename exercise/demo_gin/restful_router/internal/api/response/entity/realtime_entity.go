// @file realtime_entity.go
// @desc realtime_entity
// @author KeeneChen <keenechen@qq.com>
// @since  2024.06.25-19:50:22

package entity

type CaiyunRealtimeEntity struct {
	Status      string  `json:"status"`
	Temperature float64 `json:"temperature"` // 地表 2 米气温
	Humidity    float64 `json:"humidity"`    // 地表 2 米湿度相对湿度(%)
	Cloudrate   float64 `json:"cloudrate"`   // 总云量(0.0-1.0)
	Skycon      string  `json:"skycon"`      // 天气现象
	Visibility  float64 `json:"visibility"`  // 地表水平能见度
	Dswrf       float64 `json:"dswrf"`       // 向下短波辐射通量(W/M2)
	Wind        struct {
		Speed  float64 `json:"speed"`  // 地表 10 米风速
		Direct float64 `json:"direct"` // 地表 10 米风向
	} `json:"wind"`
	Pressure            float64 `json:"pressure"`             // 地面气压
	ApparentTemperature float64 `json:"apparent_temperature"` // 体感温度
	Precipitation       struct {
		Local struct {
			Status      string  `json:"status"`
			Datasources string  `json:"datasources"`
			Intensity   float64 `json:"intensity"` // 本地降水强度
		} `json:"precipitation"`
		Nearest struct {
			Status    string  `json:"status"`
			Distance  string  `json:"distance"`
			Intensity float64 `json:"intensity"` // 附近降水强度
		} `json:"nearest"`
	} `json:"precipitation"`
	AirQuality struct {
		Pm25 float64 `json:"pm25"` // PM2.5
		Pm10 float64 `json:"pm10"` // PM10
		O3   float64 `json:"o3"`   // 臭氧
		So2  float64 `json:"so2"`  // 二氧化硫
		No2  float64 `json:"no2"`  // 二氧化氮
		Co   float64 `json:"co"`   // 一氧化碳
		Aqi  struct {
			Chn float64 `json:"chn"` // 中国标准 AQI
			Usa float64 `json:"usa"` // 美国标准 AQI
		} `json:"aqi"`
		Description struct {
			Chn string `json:"chn"` // 中国标准 AQI 描述
			Usa string `json:"usa"` // 美国标准 AQI 描述
		} `json:"description"`
	} `json:"air_quality"`
	LifeIndex struct {
		Ultraviolet struct {
			Index float64 `json:"index"` // 紫外线指数
			Desc  string  `json:"desc"`  // 紫外线指数描述
		} `json:"ultraviolet"`
		Comfort struct {
			Index float64 `json:"index"` // 舒适度指数
			Desc  string  `json:"desc"`  // 舒适度指数描述
		} `json:"comfort"`
	} `json:"life_index"`
}
