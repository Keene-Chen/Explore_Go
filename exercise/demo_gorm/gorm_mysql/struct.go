package main

import "time"

type Person struct {
	ID        uint      `gorm:"primary_key;type:tinyint"`
	Name      string    `gorm:"type:varchar(100)"`
	Age       uint      `gorm:"type:tinyint"`
	UserName  string    `gorm:"type:varchar(100);not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Tabler interface {
	TableName() string
}

// TableName会将Person的表名重写为`person`
func (Person) TableName() string {
	return "person"
}
