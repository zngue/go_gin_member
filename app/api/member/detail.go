package member

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/common/response"
)

func Detail(ctx *gin.Context)  {

	var req  request.IDStatusRequest
	if err:=ctx.ShouldBind(&req); err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	if err:=req.CheckStatus();err!=nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	if detail,err:=Member().GetOne(req); err != nil {
		response.HttpFail(ctx,err.Error())
		return
	}else {
		response.HttpOk(ctx,detail)
		return
	}



}
