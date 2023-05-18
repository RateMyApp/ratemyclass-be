package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ratemyapp/services"
)

type courseRatingRouter struct {
	ginRouter        *gin.Engine
	courseRatingServ services.CourseRatingService
}
// handler function
func (cr *courseRatingRouter) createCourseRatingRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateCourseRatingReq

		ctx.ShouldBindJSON(&req)
		// validate request
		if validationErrorCheck(req, ctx) {
			return
		}

		command := services.CreateCourseRatingCommand{ProfessorID: req.ProfessorID, ExperienceRating: req.ExperienceRating, DifficultyRating: req.DifficultyRating, Review: req.Review, CourseID: req.CourseID, UserID: req.UserID, IsAnonymous: req.IsAnonymous}

		err := cr.courseRatingServ.CreateCourseRating(command)
		if err != nil {
			ctx.JSON(err.StatusCode, err)
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": "Course rating added"})
		return
	}
}

// add routes to router group
func (cr *courseRatingRouter) ExecRoutes() {
	routerGroup := cr.ginRouter.Group("/api/v1/courserating")
	{
		routerGroup.POST("", cr.createCourseRatingRoute())
	}
}

// constructor method 
func NewCourseRatingRouter(ginRouter *gin.Engine, courseServ services.CourseRatingService) Router {
	return &courseRatingRouter{ginRouter: ginRouter, courseRatingServ: courseServ}
}
