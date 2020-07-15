package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_member/app/api/member"
)

func MemebrRouter(group *gin.RouterGroup)  {
	api:=group.Group("member")
	{
		api.GET("list",member.List)
		api.POST("del_status",member.DelOrStatus)
		api.GET("detail",member.Detail)
		api.POST("edit",member.Edit)
	}


}
