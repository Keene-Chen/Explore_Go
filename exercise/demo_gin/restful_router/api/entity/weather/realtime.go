package entity

type CaiyunResponse struct {
	Status     string    `json:"status"`
	APIVersion string    `json:"api_version"`
	APIStatus  string    `json:"api_status"`
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	Tzshift    int       `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime int       `json:"server_time"`
	Location   []float64 `json:"location"`
	Result     Result    `json:"result"`
}

type Wind struct {
	Speed     float64 `json:"speed"`
	Direction float64 `json:"direction"`
}

type Local struct {
	Status     string  `json:"status"`
	Datasource string  `json:"datasource"`
	Intensity  float64 `json:"intensity"`
}

type Nearest struct {
	Status    string  `json:"status"`
	Distance  float64 `json:"distance"`
	Intensity float64 `json:"intensity"`
}

type Precipitation struct {
	Local   Local   `json:"local"`
	Nearest Nearest `json:"nearest"`
}

type Aqi struct {
	Chn float64 `json:"chn"`
	Usa int     `json:"usa"`
}

type Description struct {
	Usa string `json:"usa"`
	Chn string `json:"chn"`
}

type AirQuality struct {
	Pm25        float64     `json:"pm25"`
	Pm10        int         `json:"pm10"`
	O3          int         `json:"o3"`
	So2         int         `json:"so2"`
	No2         int         `json:"no2"`
	Co          float64     `json:"co"`
	Aqi         Aqi         `json:"aqi"`
	Description Description `json:"description"`
}

type Ultraviolet struct {
	Index float64 `json:"index"`
	Desc  string  `json:"desc"`
}

type Comfort struct {
	Index int    `json:"index"`
	Desc  string `json:"desc"`
}

type LifeIndex struct {
	Ultraviolet Ultraviolet `json:"ultraviolet"`
	Comfort     Comfort     `json:"comfort"`
}

type Realtime struct {
	Status              string        `json:"status"`
	Temperature         float64       `json:"temperature"`
	Humidity            float64       `json:"humidity"`
	Cloudrate           float64       `json:"cloudrate"`
	Skycon              string        `json:"skycon"`
	Visibility          float64       `json:"visibility"`
	Dswrf               float64       `json:"dswrf"`
	Wind                Wind          `json:"wind"`
	Pressure            float64       `json:"pressure"`
	ApparentTemperature float64       `json:"apparent_temperature"`
	Precipitation       Precipitation `json:"precipitation"`
	AirQuality          AirQuality    `json:"air_quality"`
	LifeIndex           LifeIndex     `json:"life_index"`
}

type Result struct {
	Realtime Realtime `json:"realtime"`
	Primary  int      `json:"primary"`
}
