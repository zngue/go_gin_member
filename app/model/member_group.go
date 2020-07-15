package model

import (
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type MemberGroup struct {
	GroupId int `gorm:"column:group_id;primary_key;type:int(11) unsigned;default:0;comment:'分组id'" json:"group_id" form:"group_id"`
	Name string `gorm:"column:name;type:varchar(50);defaulf:null;comment:'分组名称';index" json:"name" form:"name"`
	Desc string 	`gorm:"column:desc;type:varchar(100);default:null;comment:'会员分组描述'"`
	Status int 		`gorm:"column:status;type:tinyint(1);default:0;comment:'1正常 2 禁用'" json:"status" form:"status"`

}
func (MemberGroup) TableName ()  string  {

	return "member_group"
}
//添加
func (MemberGroup) Add(grooup MemberGroup) (ID int,err error) {
	err=db.MysqlConn.Create(&grooup).Error
	ID = grooup.GroupId
	return
}
//修改
func ( MemberGroup) Update(member MemberGroup ) (ID int ,err error)  {
	memberInfo:=common.StuckToMap(member)
	err=db.MysqlConn.Model(&MemberGroup{}).Update(&memberInfo).Error
	ID = member.GroupId
	return
}
//修改状态
func (MemberGroup) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&MemberGroup{}).Where( "group_id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}


