package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)
type WhereFun func(db *gorm.DB) *gorm.DB
type FindFun func(db *gorm.DB) error
type NumFun func(db *gorm.DB) error
type OrderFun func(db *gorm.DB) *gorm.DB
var DB *gorm.DB
func init()  {
	DB = db.MysqlConn
}
//统一数据处理
type Common struct {
	AddModel AddModel //增加
	UpdateModel UpdateModel  //修改
	StatusModel StatusModel //删除或者修改状态
	ListModel 	ListModel	//列表数据
	DetailModel DetailModel  //详情数据
}
type DetailModel struct {//获取详情实例
	Request  request.IDStatusRequest
	WhereFun WhereFun
	Model interface{}
	Err error
	First FindFun
}

type ListModel struct {
	Model interface{}
	Request *request.Page
	FindFun 	FindFun
	NumFun 		NumFun
	Err 		error
	WhereMap 	map[string]interface{}
	OrderFun OrderFun
}
type AddModel struct {
	Model interface{}
	Data interface{}
	AddFun FindFun
	Err error
}
type UpdateModel struct {
	Model interface{}
	Data interface{}
	Where map[string]interface{}
	WhereFun WhereFun
	Err error
}
type DelModel struct {
	Model interface{}
	Where map[string]interface{}
	Request  request.IDStatusRequest
}
type StatusModel struct {
	Request  request.IDStatusRequest
	Where map[string]interface{}
	Model interface{}
	Data interface{}
}

func (d *DetailModel) DelTail() *DetailModel {
	dbConn :=db.MysqlConn
	if d.Model!=nil {
		dbConn=dbConn.Model(&d.Model)
	}
	if d.WhereFun!=nil {
		dbConn=d.WhereFun(dbConn)
	}
	if d.Request.ID>0 {
		dbConn.Where("id = ?",d.Request.ID)
	}
	if d.First!=nil {
		d.Err=d.First(dbConn)
	}
	return d
}
//新增
func (a *AddModel) Add ( ) *AddModel {
	dbConn:=db.MysqlConn
	if a.AddFun!=nil {
		a.Err = a.AddFun(dbConn)
	}else{
		a.Err=errors.New("addfun should func")
	}
	return a
}

//修改
func (u *UpdateModel) Update() *UpdateModel  {
	commonInfo:=common.StuckToMap(u.Data)
	dbConn:=db.MysqlConn
	if u.Model!=nil{
		dbConn=dbConn.Model(&u.Model)
	}
	if u.WhereFun!=nil {
		dbConn = u.WhereFun(dbConn)
	}
	if u.Where!=nil {
		for key,val:=range u.Where {
			dbConn = dbConn.Where( key,val )
		}
	}
	u.Err = dbConn.Update(commonInfo).Error
	return u
}
//修改状态
func (d *StatusModel) UpdateStatus () (err error) {
	dbConn:=db.MysqlConn
	if d.Model!=nil {
		dbConn=dbConn.Model(&d.Model)
	}
	if d.Where!=nil {
		for key,val:=range d.Where {
			dbConn = dbConn.Where( key,val )
		}
	}else{
		if len(d.Request.IDArr)>0 {
			dbConn = dbConn.Where( "id in (?)",d.Request.IDArr  )
		}
	}
	err = dbConn.Updates(map[string]interface{}{"status":d.Request.Status}).Error
	return
}
//修改状态
func (d *StatusModel) Del () (err error) {
	dbConn:=db.MysqlConn
	if d.Model!=nil {
		dbConn=dbConn.Model(&d.Model)
	}
	if d.Where!=nil {
		for key,val:=range d.Where {
			dbConn = dbConn.Where( key,val )
		}
	}else{
		if len(d.Request.IDArr)>0 {
			dbConn = dbConn.Where( "id in (?)",d.Request.IDArr  )
		}
	}
	err=dbConn.Delete(&d.Model).Error
	return
}
//获取列表数据
func (l *ListModel) List(listFun  ...WhereFun) *ListModel {
	dbConn := db.MysqlConn
	if l.Model!=nil {
		dbConn=dbConn.Model(&l.Model)
	}
	if len(listFun)>0 {//参数where条件
		for _,fn :=range listFun{
			dbConn=fn(dbConn)
		}
	}
	if l.WhereMap!=nil {//默认where条件
		for key,val:=range l.WhereMap {
			dbConn = dbConn.Where( key,val )
		}
	}
	if l.NumFun!=nil && l.Request.IsCount!=0{//返回列表数据
		err :=l.NumFun(dbConn)
		if err!=nil {
			l.Err=err
			return l
		}
		if l.Request.IsCount==2 {
			return l
		}
	}
	if l.OrderFun!=nil {
		dbConn = l.OrderFun(dbConn)
	}
	if l.Request!=nil && l.Request.Page >0 && l.Request.PageSize>0{
		dbConn=dbConn.Offset(l.Request.PageLimitOffset()).Limit(l.Request.PageSize)
	}
	if l.FindFun!=nil  {
		l.Err=l.FindFun(dbConn)
		return l
	}
	return l
}
