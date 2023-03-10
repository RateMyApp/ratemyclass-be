package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ratemyapp/dto/requests"
)

type HelloRouter struct {
	ginRouter *gin.Engine
}

func (hr *HelloRouter) sayHelloRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request requests.SayHelloRequest

		ctx.ShouldBindJSON(&request)

		if requestErr := request.Validate(); requestErr != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": requestErr})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
	}
}

func (hr *HelloRouter) ExecRoutes() {
	routerGroup := hr.ginRouter.RouterGroup
	{
		routerGroup.GET("/HELLO", hr.sayHelloRoute())
	}
}

func NewHelloRouter(ginRouter *gin.Engine) Router {
	return &HelloRouter{ginRouter: ginRouter}
}
