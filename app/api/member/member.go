package member

import "github.com/zngue/go_gin_member/app/service/member"

func Member() *member.Member  {
	return new(member.Member)
}
