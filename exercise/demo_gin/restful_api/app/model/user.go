package model

type User struct {
	// gorm.Model
	Id   int    `json:"userId" gorm:"primaryKey;autoIncrement;"` //对应表中的字段名 id
	Name string `json:"name" gorm:"size:30;"`                    //name
	Sex  int    `json:"sex" gorm:"size:1;"`                      //sex
}
