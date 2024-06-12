package model

import (
	"fmt"
	"restful_api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 取得数据库连接实例
func DB() *gorm.DB {
	var config config.Config
	conf := config.GetConf()
	fmt.Println("model.DB().conf.Database.Source=", conf.Database.Source)
	dsn := conf.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "kc_",
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("err: ", err)
	}
	return db
}
