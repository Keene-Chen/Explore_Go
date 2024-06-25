package response

import "Explore_Go/exercise/demo_gin/restful_router/internal/api/response/entity"

type CaiyunRealtimeAlertResponse struct {
	entity.CaiyunCommonEntity
	Result struct {
		Alert    entity.CaiyunAlertEntity `json:"alert"`
		Realtime CaiyunRealtimeResponse   `json:"realtime"`
		Primary  int                      `json:"primary"` // 主要预警
	} `json:"result"`
}

type CaiyunMinutelyAlertResponse struct {
	entity.CaiyunCommonEntity
	Result struct {
		Alert    entity.CaiyunAlertEntity `json:"alert"`
		Minutely CaiyunMinutelyResponse   `json:"minutely"`
		Primary  int                      `json:"primary"` // 主要预警
	} `json:"result"`
}

type CaiyunHourlyAlertResponse struct {
	entity.CaiyunCommonEntity
	Result struct {
		Alert   entity.CaiyunAlertEntity `json:"alert"`
		Hourly  CaiyunHourlyResponse     `json:"hourly"`
		Primary int                      `json:"primary"` // 主要预警
	} `json:"result"`
}

type CaiyunDailyAlertResponse struct {
	entity.CaiyunCommonEntity
	Result struct {
		Alert   entity.CaiyunAlertEntity `json:"alert"`
		Daily   CaiyunDailyResponse      `json:"daily"`
		Primary int                      `json:"primary"` // 主要预警
	} `json:"result"`
}
