package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
}

func ReadConfig() []Config {
	// 读取配置文件
	data, err := os.ReadFile("config.json")
	if err != nil {
		panic("failed to read config file")
	}

	var config []Config
	json.Unmarshal(data, &config)
	return config
}

func ConnectGorm() (db *gorm.DB, err error) {
	config := ReadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config[0].UserName, config[0].Password, config[0].Host, config[0].Port, config[0].DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 测试连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database successfully!")

	return db, err
}
