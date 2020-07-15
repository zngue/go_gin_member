package member

import (
	"github.com/zngue/go_gin_member/app/model"
	"github.com/zngue/go_gin_member/app/request/memeber"
)

type ReturnData struct {
	Count int
	Item []*model.Members
}
func (s *Member) List ( request memeber.ListRequest ) (num int,mb model.Members,err error ) {
	dbConn:=s.MemberModel.GetList()
	if len(request.IDArr)>0 {
		dbConn=dbConn.Where(" id in (?) ",request.IDArr)
	}
	if request.Name!="" {
		dbConn = dbConn.Where( "name like ?","%"+request.Name+"%" )
	}
	if request.Mobile!=0 {
		dbConn = dbConn.Where("mobile = ?",request.Mobile)
	}
	if request.IsCount!=0  {
		err=dbConn.Count(&num).Error
		if err!=nil {
			return
		}
		if request.IsCount==2 {
			return
		}
	}
	if request.PageSize>0 && request.Page.Page>0 {
		dbConn = dbConn.Offset(request.PageLimitOffset()).Limit(request.PageSize)
	}
	dbConn = dbConn.Order("id desc")
	err =dbConn.Find(&mb).Error
	if err != nil {
		return
	}
	return




}
