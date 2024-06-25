package response

import "Explore_Go/exercise/demo_gin/restful_router/internal/api/response/entity"

type CaiyunRealtimeResponse struct {
	entity.CaiyunCommonEntity
	Result struct {
		Realtime entity.CaiyunRealtimeEntity `json:"realtime"`
		Primary  int                         `json:"primary"`
	} `json:"result"`
}

type CaiyunDailyResponse struct {
	entity.CaiyunCommonEntity
	Result struct {
		Daily   entity.CaiyunDailyEntity `json:"daily"`
		Primary int                      `json:"primary"` // 主要预报
	} `json:"result"`
}

type CaiyunHourlyResponse struct {
	entity.CaiyunCommonEntity
	Result struct {
		Hourly           entity.CaiyunHourlyEntity `json:"hourly"`
		Primary          int                       `json:"primary"`           // 主要预报
		ForecastKeyPoint string                    `json:"forecast_keypoint"` // 预报关键点
	} `json:"result"`
}

type CaiyunMinutelyResponse struct {
	entity.CaiyunCommonEntity
	Result struct {
		Minutely         entity.CaiyunMinutelyEntity `json:"minutely"`
		Primary          int                         `json:"primary"`           // 主要预报
		ForecastKeypoint string                      `json:"forecast_keypoint"` // 预报关键点
	} `json:"result"`
}
