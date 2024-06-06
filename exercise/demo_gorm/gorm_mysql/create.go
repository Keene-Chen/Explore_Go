package main

import (
	"crypto/sha256"
	"fmt"

	"gorm.io/gorm"
)

func CreateOne(db *gorm.DB) {
	// 创建记录,对密码字段进行加密
	h := sha256.Sum256([]byte("123456"))
	hashPassword := fmt.Sprintf("%x", h)
	db.Create(&Person{Name: "Keene", Age: 18, UserName: "Keene", Password: hashPassword})
}
