package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ratemyapp/services"
)

type professorRouter struct {
	ginRouter     *gin.Engine
	professorServ services.ProfessorService
}

func (pr *professorRouter) createProfRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateProfReq

		ctx.ShouldBindJSON(&req)

		if validationErrorCheck(req, ctx) {
			return
		}

		command := services.CreateProfessorCommand{Email: req.Email, Department: req.Department, DirectoryListing: req.DirectoryListing, Firstname: req.Firstname, Lastname: req.Lastname}

		err := pr.professorServ.CreateProfessor(command)
		if err != nil {
			ctx.JSON(err.StatusCode, err)
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": "Professor created"})
		return
	}
}

func (pr *professorRouter) ExecRoutes() {
	routerGroup := pr.ginRouter.Group("/api/v1/prof")

	{
		routerGroup.POST("", pr.createProfRoute())
	}
}

func NewProfessorRouter(ginRouter *gin.Engine, professorServ services.ProfessorService) Router {
	return &professorRouter{ginRouter: ginRouter, professorServ: professorServ}
}
