package routers

import (
	// "encoding/json"

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
)

func validationErrorCheck(req *validation.Validatable, ctx *gin.Context) bool {
	err := ctx.ShouldBindJSON(&req)
	err := dto.Validate()
	if err != nil {
		appError := exceptions.NewBadRequestError(err)
		ctx.JSON(appError.StatusCode, appError)
		return true
	}
	return false
}
