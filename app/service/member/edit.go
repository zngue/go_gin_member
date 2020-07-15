package member
import "github.com/zngue/go_gin_member/app/model"
func (m *Member) Edit(members model.Members) (int ,error) {
	if members.ID>0 {
		return m.MemberModel.Update(members)
	}else {
		return m.MemberModel.Add(members)
	}
}
