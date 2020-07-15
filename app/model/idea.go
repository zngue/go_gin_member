package model

import "github.com/zngue/go_tool/src/db"

type Idea struct {
	
	ID int `gorm:"column:id;primary;type:int(11);" json:"id" form:"id"`
	Title string `gorm:"column:title;index;type:varchar(255);defaulf:null" json:"title" form:"title"`
	Keywords string `gorm:"column:keywords;type:varchar(255);defaulf:null" json:"keywords" form:"keywords"`
	Description string `gorm:"column:description;type:varchar(255);defaulf:null" json:"description" form:"description"`
	Content string 	`gorm:"column:content;type:text;defaulf:null" json:"content" form:"content"`
	CodeOne string `gorm:"column:code_one;type:text;defaulf:null"  form:"code_one" json:"code_one"`
	CodeTwo string `gorm:"column:code_two;type:text;defaulf:null"  form:"code_two" json:"code_two"`
	CateID int 		`gorm:"column:cate_id;type:int(10)" json:"cate_id"`
	Cate  []Cate		`json:"cate" gorm:"many2many:idea_tag"`
	db.TimeStampModel
}

func (Idea) TableName() string {
	return "idea"
}