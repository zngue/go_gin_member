package model

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type Members struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	OpenID     string `gorm:"index;column:openid;type:varchar(50)" json:"openid"`
	Password string		`gorm:"column:password;type:varchar(100)"`
	UnionID    string `gorm:"index;column:unionid;type:varchar(50)" json:"unionid"`
	Mobile     int	  `gorm:"index;column:mobile;" json:"mobile"`
	Credits    int    `gorm:"column:credits;type:int(10)" json:"credits"`
	Channel    int    `gorm:"column:channel;type:int(10)" json:"channel"`
	Avatar     string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
	GroupID    int    `gorm:"index;column:group_id;type:int(10);" json:"group_id"`
	Name       string `gorm:"index;column:name;type:varchar(50);" json:"name"`
	Gender     int8   `gorm:"column:gender;type:tinyint(1);" json:"gender"`
	Address    string `gorm:"column:address;type:varchar(200);" json:"address"`
	Status     int8   `gorm:"index;column:status;type:tinyint(1)" json:"status"`
	db.TimeStampModel
}

func (Members) TableName() string  {
	return "members"
}
//添加
func (Members) Add(Members Members) (ID int,err error) {
	err=db.MysqlConn.Create(&Members).Error
	ID = Members.ID
	return
}
//修改
func ( Members) Update(member Members ) (ID int ,err error)  {
	memberInfo:=common.StuckToMap(member)
	err=db.MysqlConn.Model(&Members{}).Update(&memberInfo).Error
	ID = member.ID
	return
}
//修改状态
func (Members) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&Members{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (Members) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&Members{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&Members{}).Error
	return
}
func (Members) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *Members) GetOne (request request.IDStatusRequest ) ( mb *Members,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






