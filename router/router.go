package router

import "github.com/gin-gonic/gin"

func Router(group *gin.RouterGroup)  {
	MemebrRouter(group)
	IdeaRouter(group)

}
