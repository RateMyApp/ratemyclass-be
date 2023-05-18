package routers

import "github.com/gin-gonic/gin"

type courseRouter struct {
	ginRouter *gin.Engine
}

func (self *courseRouter) createCourse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
