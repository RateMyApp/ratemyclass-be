package routers

import (
	// "encoding/json"

	"log"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/ratemyapp/exceptions"
	"go.uber.org/fx"
)

// interface method implemented by all routes, this adds the route, and handler function to the router group
type Router interface {
	ExecRoutes()
}

// helper function
func AsRouter(f interface{}) any {
	return fx.Annotate(
		f,
		fx.As(new(Router)),
		fx.ResultTags(`group:"routers"`),
	)
}

// create a module, and instantiate our newhelloRouter and NewbyeRouter objects, resolve all depedencies needed by both objects
var Module = fx.Module("routers",
	fx.Provide(AsRouter(NewHelloRouter)),
	fx.Provide(AsRouter(NewByeRouter)),
	fx.Provide(AsRouter(NewAuthRouter)),
	fx.Provide(AsRouter(NewProfessorRouter)),
	fx.Provide(AsRouter(NewCourseRouter)),
)

// Requires an address to a request dto (e.g validationErrorCheck(&req)) and checks the following:
// 1. If the request body (in json) can be parsed into that dto
// 2. If the request dto has the Validate() method then validation will be executed
// If an error occurs in any of these checks then the method returns true and the gin context
// is populated with the appropriate message and status code
func validationErrorCheck(req interface{}, ctx *gin.Context) bool {
	var appError exceptions.AppError

	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Println(err)
		appError = exceptions.NewUnprocessableEntityError("Error reading json")
		ctx.JSON(appError.StatusCode, appError)
		return true
	}

	val, ok := req.(validation.Validatable)

	if ok {
		if err := val.Validate(); err != nil {
			appError = exceptions.NewBadRequestError(err)
			ctx.JSON(appError.StatusCode, appError)
			return true
		}
	}

	return false
}
