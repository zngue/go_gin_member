package member

import (
	"errors"
	"github.com/zngue/go_tool/src/common/request"
)
func (m *Member) DelOrStatus(request request.IDStatusRequest) error {
	if len(request.IDArr)==0 {
		return errors.New("参数错误")
	}
	if request.From==1 {
		return m.MemberModel.UpdateStatus(request)
	}else{
		return m.MemberModel.Del(request)
	}
}
