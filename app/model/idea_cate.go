package model

import "github.com/zngue/go_tool/src/db"

type Cate struct {
	ID int 		`gorm:"column:id;primary_key;auto_increment"`
	Name string `gorm:"column:name;type:varchar(255)"`
	Desc string `gorm:"column:desc;type:varchar(255)"`
	Sort int 	`gorm:"column:sort;type:int(11)"`
	Status int 	`gorm:"column:status;type:tinyint(1)"`
	Keywords string `gorm:"column:keywords;type:varchar(255);defaulf:null" json:"keywords" form:"keywords"`
	Description string `gorm:"column:description;type:varchar(255);defaulf:null" json:"description" form:"description"`
	Pid int		`gorm:"column:pid;type:int(11)"`
	db.TimeStampModel
}

func (Cate) TableName() string  {
	return "idea_cate"
}