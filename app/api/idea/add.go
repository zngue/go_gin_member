package idea

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_member/app/model"
	"github.com/zngue/go_gin_member/app/service/idea"
	"github.com/zngue/go_gin_member/app/service/idea_tag"
	"github.com/zngue/go_tool/src/common/response"
	"github.com/zngue/go_tool/src/fun/time"
	"sync"
)
var IdeaListView []*model.Cate

func Add (c *gin.Context)  {

	list,err:=IdeaCate().List()
	IdeaListView = list
	if err!=nil {
		response.HttpFail(c)
	}
	ideas :=new(idea.Ideas).GetAllIdea()
	model:=new(idea.Idea)
	var wg sync.WaitGroup
	for _,cateOne:=range list{
		wg.Add(1)
		go IdeaAdd(cateOne,ideas,model ,&wg ,IdeaListView)
	}
	wg.Wait()

	response.HttpOk(c)
	return
}
func IdeaAdd(cate *model.Cate,ideas *idea.Ideas,idea *idea.Idea,wg *sync.WaitGroup,IdeaListView []*model.Cate)  {
	defer func() {
		if r:=recover(); r != nil {
			panic(r)
			fmt.Println(r)
		}
	}()
	defer wg.Done()
	models :=model.Idea{
		Title: "2"+time.TimeToFormat(time.Time(), "20060102") + cate.Name+"激活码，"+ cate.Name+"免费注册码,"+ cate.Name+" 最新激活码",
		Keywords: cate.Keywords,
		Description:cate.Description,
		Content: time.TimeToFormat(time.Time(), "2006-01-02 15:04:05更新")+cate.Name+"激活码，"+cate.Name+"免费注册码",
		CodeOne: ideas.One,
		CodeTwo: ideas.Two,
		CateID: cate.ID,
	}
	id,err:=idea.Edit(models)
	if err==nil && id>0 && cate.ID>0 {
		tagModel:=	new(idea_tag.IdeaTag)
		for _,ca:= range IdeaListView{
			tag:=model.IdeaTag{
				IdeaID: id,
				CateID: ca.ID,
			}
			tagModel.Add(tag)
		}
	}

}
