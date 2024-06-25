// @file daily_entity.go
// @desc daily_entity
// @author KeeneChen <keenechen@qq.com>
// @since  2024.06.25-19:50:00

package entity

type DailyPrecipitation struct {
	Date        string  `json:"date"`
	Max         float64 `json:"max"`         // 最大降水量
	Min         float64 `json:"min"`         // 最小降水量
	Avg         float64 `json:"avg"`         // 平均降水量
	Probability float64 `json:"probability"` // 降水概率
}

type DailyTemperature struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"` // 最高温度
	Min  float64 `json:"min"` // 最低温度
	Avg  float64 `json:"avg"` // 平均温度
}

type DailyWind struct {
	Date string `json:"date"`
	Max  struct {
		Speed     float64 `json:"speed"`     // 最大风速
		Direction float64 `json:"direction"` // 最大风向
	} `json:"max"`
	Min struct {
		Speed     float64 `json:"speed"`     // 最小风速
		Direction float64 `json:"direction"` // 最小风向
	} `json:"min"`
	Avg struct {
		Speed     float64 `json:"speed"`     // 平均风速
		Direction float64 `json:"direction"` // 平均风向
	} `json:"avg"`
}

type DailySkycon struct {
	Date  string `json:"date"`
	Value string `json:"value"` // 天气现象
}

type DailyLifeIndexItem struct {
	Date  string `json:"date"`
	Index string `json:"index"` // 指数
	Desc  string `json:"desc"`  // 指数描述
}

type CaiyunDailyEntity struct {
	Status string `json:"status"` // "ok" 天级预报状态
	Astro  []struct {
		Date    string `json:"date"`
		Sunrise struct {
			Time string `json:"time"` // 日出时间
		} `json:"sunrise"`
		Sunset struct {
			Time string `json:"time"` // 日落时间
		} `json:"sunset"`
	} `json:"astro"` // 天文数值
	Precipitation_08h_20h []DailyPrecipitation `json:"precipitation_08h_20h"` // 白天降水数据
	Precipitation_20h_32h []DailyPrecipitation `json:"precipitation_20h_08h"` // 夜晚降水数据
	Precipitation         []DailyPrecipitation `json:"precipitation"`         // 24 小时降水数据
	Temperature_08h_20h   []DailyTemperature   `json:"temperature_08h_20h"`   // 白天温度数据
	Temperature_20h_32h   []DailyTemperature   `json:"temperature_20h_32h"`   // 夜晚温度数据
	Temperature           []DailyTemperature   `json:"temperature"`           // 24 小时温度数据
	Wind_08h_20h          []DailyWind          `json:"wind_08h_20h"`          // 白天风力数据
	Wind_20h_32h          []DailyWind          `json:"wind_20h_32h"`          // 夜晚风力数据
	Wind                  []DailyWind          `json:"wind"`                  // 24 小时风力数据
	Humidity              []DailyTemperature   `json:"humidity"`              // 地表 2 米相对湿度(%)
	Cloudrate             []DailyTemperature   `json:"cloudrate"`             // 云量(0.0-1.0)
	Pressure              []DailyTemperature   `json:"pressure"`              // 气压
	Visibility            []DailyTemperature   `json:"visibility"`            // 能见度(km)
	Dswrf                 []DailyTemperature   `json:"dswrf"`                 // 全天累计直射辐射量(W/m^2)
	AirQuality            []struct {
		Aqi []struct {
			Date string `json:"date"`
			Max  struct {
				Chn float64 `json:"chn"` // 国内 AQI
				Usa int     `json:"usa"` // 美国 AQI
			} `json:"max"` // 最大空气质量指数
			Min struct {
				Chn float64 `json:"chn"` // 国内 AQI
				Usa int     `json:"usa"` // 美国 AQI
			} `json:"min"` // 最小空气质量指数
			Avg struct {
				Chn float64 `json:"chn"` // 国内 AQI
				Usa int     `json:"usa"` // 美国 AQI
			} `json:"avg"` // 平均空气质量指数
		} `json:"aqi"`
		Pm25 []DailyTemperature `json:"pm25"` // PM2.5
	} `json:"air_quality"` // 空气质量
	Skycon         []DailySkycon `json:"skycon"`         // 天气现象
	Skycon_08h_20h []DailySkycon `json:"skycon_08h_20h"` // 白天天气现象
	Skycon_20h_32h []DailySkycon `json:"skycon_20h_32h"` // 夜晚天气现象
	LifeIndex      struct {
		Ultraviolet []DailyLifeIndexItem `json:"ultraviolet"` // 紫外线指数
		CarWashing  []DailyLifeIndexItem `json:"car_washing"` // 洗车指数
		Dressing    []DailyLifeIndexItem `json:"dressing"`    // 穿衣指数
		Comfort     []DailyLifeIndexItem `json:"comfort"`     // 舒适度指数
		ColdRisk    []DailyLifeIndexItem `json:"cold_risk"`   // 感冒指数
	} `json:"life_index"` // 生活指数
}
