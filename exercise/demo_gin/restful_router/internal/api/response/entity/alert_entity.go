// @file alert_entity.go
// @desc alert_entity
// @author KeeneChen <keenechen@qq.com>
// @since  2024.06.25-19:49:33

package entity

type CaiyunAlertEntity struct {
	Status  string `json:"status"` // "ok" 预警状态
	Content []struct {
		Province      string    `json:"province"`      // 省份
		Status        string    `json:"status"`        // 预警状态
		Code          string    `json:"code"`          // 预警代码
		Description   string    `json:"description"`   // 预警描述
		RegionIds     []string  `json:"regionId"`      // 预警区域 ID
		County        string    `json:"county"`        // 县
		Pubtimestamp  int       `json:"pubtimestamp"`  // 预警发布时间
		Latlon        []float64 `json:"latlon"`        // 预警区域经纬度
		City          string    `json:"city"`          // 城市
		AlertId       string    `json:"alertId"`       // 预警 ID
		Title         string    `json:"title"`         // 预警标题
		Adcode        string    `json:"adcode"`        // 预警区域行政区划代码
		Source        string    `json:"source"`        // 预警来源
		Location      string    `json:"location"`      // 预警区域名称
		RequestStatus string    `json:"requestStatus"` // 请求状态
	} `json:"content"` // 预警内容
	Adcodes []struct {
		Adcode string `json:"adcode"` // 行政区划代码
		Name   string `json:"name"`   // 行政区划名称
	} `json:"adcodes"` // 行政区划代码
}
