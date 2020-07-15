package member

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_member/app/model"
	"github.com/zngue/go_tool/src/common/response"
)
//编辑
func Edit(ctx *gin.Context)  {
	var mb model.Members
	if err:=ctx.ShouldBind(&mb); err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	if id,err:=Member().Edit(mb); err != nil {
		response.HttpFail(ctx,err.Error())
		return
	}else {
		response.HttpFail(ctx,id)
		return
	}



}
