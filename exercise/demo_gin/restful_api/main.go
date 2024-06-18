package restful_api

import (
	"Explore_Go/exercise/demo_gin/restful_api/config"
	"Explore_Go/exercise/demo_gin/restful_api/router"
	"encoding/json"
	"fmt"
)

func main() {
	//加载配置
	var config config.Config
	conf := config.GetConf()

	//将对象，转换成json格式
	data_config, err := json.Marshal(conf)

	if err != nil {
		fmt.Println("err:\t", err.Error())
		return
	}
	fmt.Println("data_config:\t", string(data_config))
	fmt.Println("config.Database.Driver=", config.Database.Driver)

	//加载路由
	r := router.SetupRouter()
	err = r.Run("127.0.0.1:8080")
	if err != nil {
		fmt.Println("err:\t", err.Error())
		return
	}
}
