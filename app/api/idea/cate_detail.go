package idea

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/common/response"
)

func CateDetail(ctx *gin.Context)  {
	var req request.IDStatusRequest
	if err:=ctx.ShouldBind(&req); err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	if detail,err:=IdeaCate().Detail(req); err != nil {
		response.HttpFail(ctx,err.Error())
		return
	}else {
		response.HttpOk(ctx,detail)
		return
	}


}


