package member

import (
	"github.com/zngue/go_gin_member/app/model"
	"github.com/zngue/go_tool/src/common/request"
)

func (m *Member) GetOne(request request.IDStatusRequest) (mb *model.Members,err error)  {
	return m.MemberModel.GetOne(request)
}
