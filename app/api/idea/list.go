package idea

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_member/app/request/idea"
	"github.com/zngue/go_tool/src/common/response"
)

func List(ctx *gin.Context)  {
	var request	idea.ListRequest
	if err:=ctx.ShouldBind(&request); err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	request.IsCount=0
	if list ,err:=Idea().List(request); err != nil {
		response.HttpFail(ctx,err.Error())
		return
	}else {
		response.HttpOk(ctx,list)
		return
	}
}
