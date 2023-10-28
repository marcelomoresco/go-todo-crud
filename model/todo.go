package model

type Todo struct {
	Id int `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Done bool `gorm:"type:boolean"`
}