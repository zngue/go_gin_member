package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_member/app/api/idea"
)

func IdeaRouter(group *gin.RouterGroup)  {

	ideas :=group.Group("idea")
	{
		ideas.GET("detail",idea.Detail)
		ideas.GET("list",idea.List)
		ideas.GET("uadd",idea.Add)

		ideas.GET("cate_list",idea.CateList)
		ideas.GET("cate_detail",idea.CateDetail)
	}

}
