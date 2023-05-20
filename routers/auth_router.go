package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ratemyapp/services"
)

type authRouter struct {
	ginRouter   *gin.Engine
	authService services.AuthService
}

// register Route
func (ar *authRouter) registerRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req registerUserReq

		// validation
		if validationErrorCheck(&req, ctx) {
			return
		}

		// register command
		var command services.RegisterCommand = services.RegisterCommand{Email: req.Email, Password: req.Password, Firstname: req.Firstname, Lastname: req.Lastname}
		err := ar.authService.RegisterUser(command)
		if err != nil {
			ctx.JSON(err.StatusCode, err)
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": "User Registration Successful!!!"})
	}
}

// login route
func (ar *authRouter) loginRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req loginUserReq

		// validation
		if validationErrorCheck(&req, ctx) {
			return
		}

		var command services.LoginCommand = services.LoginCommand{Email: req.Email, Password: req.Password}

		err, user := ar.authService.LoginUser(command)
		// error found
		if err != nil {
			ctx.JSON(err.StatusCode, err)
			return
		}

		resp := loginUserResp{
			Email:       user.Email,
			Firstname:   user.Firstname,
			Lastname:    user.Lastname,
			AccessToken: user.AccessToken,
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (ar *authRouter) ExecRoutes() {
	routerGroup := ar.ginRouter.Group("/api/v1/auth")
	{
		routerGroup.POST("/register", ar.registerRoute())
		routerGroup.POST("/login", ar.loginRoute())
	}
}

// Initializes a new Router responsible for authentication
func NewAuthRouter(ginRouter *gin.Engine, authService services.AuthService) Router {
	return &authRouter{ginRouter, authService}
}
