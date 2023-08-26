package routers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ratemyapp/services"
)

type schoolRouter struct {
	router        *gin.Engine
	schoolService services.SchoolService
}

func (sr *schoolRouter) CreateSchool() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateSchoolReq

		if validationErrorCheck(&req, ctx) {
			return
		}

		command := services.CreateSchoolCommand{Name: req.Name, Address: req.Address, City: req.City, ProvinceOrState: req.PronviceOrState, Country: req.Country, EmailDomains: req.EmailDomains}

		err := sr.schoolService.CreateSchool(command)

		if err != nil {
			ctx.JSON(err.StatusCode, err)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "School has beed created"})
	}
}

func (sr *schoolRouter) SearchForSchool() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.DefaultQuery("search", "")

		query = strings.TrimSpace(query)
		res := make([]SearchForSchoolInfoRes, 0)

		if query == "" {
			ctx.JSON(http.StatusOK, res)
		}

		schools, err := sr.schoolService.SearchSchoolsInfoByName(query)

		if err != nil {
			ctx.JSON(err.StatusCode, err)
			return
		}

		for _, val := range schools {
			res = append(res, SearchForSchoolInfoRes{Name: val.Name, Location: val.Location, Id: val.Id})
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func (sr *schoolRouter) ExecRoutes() {
	routerGroup := sr.router.Group("/api/v1/school")
	{
		routerGroup.POST("/", sr.CreateSchool())
		routerGroup.GET("", sr.SearchForSchool())
	}
}

func NewSchoolRouter(router *gin.Engine, schoolService services.SchoolService) Router {
	return &schoolRouter{router: router, schoolService: schoolService}
}
