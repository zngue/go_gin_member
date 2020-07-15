package member

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/common/response"
)

func DelOrStatus(ctx *gin.Context)  {
	var req  request.IDStatusRequest
	if err:=ctx.ShouldBind(&req); err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	if errs :=req.CheckStatus(); errs != nil {
		response.HttpFailWithParameter(ctx,errs.Error())
		return
	}
	if qer:=Member().DelOrStatus(req); qer != nil {
		response.HttpFail(ctx,qer.Error())
		return
	}
	response.HttpOk(ctx)
	return
}