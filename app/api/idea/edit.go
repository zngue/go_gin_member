package idea

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_member/app/model"
	"github.com/zngue/go_tool/src/common/response"
)

func Edit(ctx *gin.Context)  {
	var idea  model.Idea
	if err:=ctx.ShouldBind(&idea); err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
}
