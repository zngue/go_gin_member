package idea

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_tool/src/common/response"
)

func CateList(ctx *gin.Context)  {
	if list,err :=IdeaCate().List(); err != nil {
		response.HttpFail(ctx,err.Error())
		return
	}else {
		response.HttpOk(ctx,list)
	}


}
