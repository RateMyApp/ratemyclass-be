package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ratemyapp/services"
)

type courseRouter struct {
	ginRouter     *gin.Engine
	courseService services.CourseService
	}

func (self *courseRouter) createCourse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createCourseReq

		if validationErrorCheck(&req, ctx) {
			return
		}

		command := services.CreateCourseCommand{
			Code:  req.Code,
			Name:  req.Name,
			Units: req.Units,
		}

		err := self.courseService.CreateCourse(command)
		if err != nil {
			ctx.JSON(err.StatusCode, err)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Course has been submitted"})
	}
}

func (self *courseRouter) ExecRoutes() {
	routerGroup := self.ginRouter.Group("/api/v1/course")
	{
		routerGroup.POST("", self.createCourse())
	}
}

func NewCourseRouter(ginRouter *gin.Engine, courseService services.CourseService) Router {
	return &courseRouter{
		ginRouter:     ginRouter,
		courseService: courseService,
	}
}
