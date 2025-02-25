package infrastructure

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	routes := router.Group("pull_request")
	{
		routes.POST("/process", HandlePullRequestEvent)
	}

	router.POST("/", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/pull_request/process"
		router.HandleContext(ctx)
	})
}
