// @file minutely_entity.go
// @desc minutely_entity
// @author KeeneChen <keenechen@qq.com>
// @since  2024.06.25-19:50:13

package entity

type CaiyunMinutelyEntity struct {
	Status          string    `json:"status"`           // "ok" 分钟级预报状态
	DataSource      string    `json:"datasource"`       // 数据源
	Precipitation2h []float64 `json:"precipitation_2h"` // 未来2小时每分钟的雷达降水强度
	Precipitation   []float64 `json:"precipitation"`    // 未来1小时每分钟的雷达降水强度
	Probability     []float64 `json:"probability"`      // 未来两小时每半小时的降水概率
	Description     string    `json:"description"`      // 预报描述
}
