package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ByeRouter struct {
	ginRouter *gin.Engine
}

func (hr *ByeRouter) sayHelloRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
	}
}

func (hr *ByeRouter) ExecRoutes() {
	routerGroup := hr.ginRouter.RouterGroup
	{
		routerGroup.GET("/bye", hr.sayHelloRoute())
	}
}

func NewByeRouter(ginRouter *gin.Engine) Router {
	return &ByeRouter{ginRouter: ginRouter}
}
