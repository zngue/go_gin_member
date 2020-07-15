package member

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_member/app/request/memeber"
	"github.com/zngue/go_tool/src/common/response"
)

func List(ctx *gin.Context)  {

	var req  memeber.ListRequest
	if err:=ctx.ShouldBind(&req); err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	num,list,errs :=Member().List(req)

	if errs!=nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx,gin.H{
		"count":num,
		"item":list,
	})
	return
}
