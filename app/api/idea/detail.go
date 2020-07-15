package idea

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/common/response"
)

func Detail(ctx *gin.Context)  {
	var request request.IDStatusRequest
	if err:=ctx.ShouldBind(&request); err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	if errss:=request.CheckStatus(); errss != nil {
		response.HttpFailWithParameter(ctx,errss.Error())
		return
	}
	if detail,errs:=Idea().Detail(request); errs != nil {
		response.HttpFail(ctx,errs.Error())
		return

	}else{
		response.HttpOk(ctx,detail)
		return
	}

}
