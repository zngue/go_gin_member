package idea

import (
	"github.com/zngue/go_gin_member/app/service/idea"
	"github.com/zngue/go_gin_member/app/service/idea_cate"
)

func Idea()  *idea.Idea {
	return new(idea.Idea)
}
func IdeaCate ()  *idea_cate.IdeaCate {
	return new(idea_cate.IdeaCate)
}
