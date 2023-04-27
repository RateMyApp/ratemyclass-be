package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ByeRouter struct {
	ginRouter *gin.Engine
}

// handler function
func (hr *ByeRouter) sayHelloRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
	}
}

// add route, and handler function to router group
func (hr *ByeRouter) ExecRoutes() {
	routerGroup := hr.ginRouter.RouterGroup
	{
		routerGroup.GET("/bye", hr.sayHelloRoute())
	}
}

// factory method/constructor to create Byerouter, dependency is a ginEngine
func NewByeRouter(ginRouter *gin.Engine) Router {
	return &ByeRouter{ginRouter: ginRouter}
}
