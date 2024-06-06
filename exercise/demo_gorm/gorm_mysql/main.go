package main

import (
	"fmt"
)

func main() {
	db, err := ConnectGorm()
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移
	// err = db.AutoMigrate(&Person{})
	// if err != nil {
	// 	panic("failed to migrate")
	// }

	// 创建记录,对密码字段进行加密
	// h := sha256.Sum256([]byte("123456"))
	// hashPassword := fmt.Sprintf("%x", h)
	// db.Create(&Person{Name: "Keene", Age: 18, UserName: "Keene", Password: hashPassword})

	// 创建多条记录
	// user := []*Person{
	// 	{Name: "K7N", Age: 19},
	// 	{Name: "KC", Age: 20},
	// }
	// db.Create(user)

	// 查询
	var person Person
	db.First(&person, 1) // 查询ID为1的person
	fmt.Println(person)

	// 查询所有
	var persons []Person
	db.Find(&persons)
	fmt.Println(persons)
}
