// @file common_entity.go
// @desc common_entity
// @author KeeneChen <keenechen@qq.com>
// @since  2024.06.25-19:49:41

package entity

type CaiyunCommonEntity struct {
	Status     string    `json:"status"`
	APIVersion string    `json:"api_version"`
	APIStatus  string    `json:"api_status"`
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	Tzshift    int       `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime int       `json:"server_time"`
	Location   []float64 `json:"location"`
}
