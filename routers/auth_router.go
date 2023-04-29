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
		var req signUpDto

		ctx.ShouldBindJSON(&req)

		// validation
		if validationErrorCheck(req, ctx) {
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
//login route
func(ar *authRouter) loginRoute() gin.HandlerFunc{
	return func(ctx *gin.Context){
		var req loginDto

		ctx.BindJSON(&req)

		var command services.LoginCommand = services.LoginCommand{Email: req.Email, Password: req.Password}
		
		err,user:= ar.authService.CheckLogin(command)
		//error found
		if err!= nil{
			ctx.JSON(err.StatusCode,err)
		}
		var resp services.UserDetails
		//login unsuccesful, empty struct
		if user ==resp{
			ctx.JSON(http.StatusOK,gin.H{"message:": "Login Unsuccessful"})
		}else{
			//create jwtToken
		
		}
	}
}
func (ar *authRouter) ExecRoutes() {
	routerGroup := ar.ginRouter.Group("/api/v1/auth")
	{
		routerGroup.POST("/register", ar.registerRoute())
	}
}

func NewAuthRouter(ginRouter *gin.Engine, authService services.AuthService) Router {
	return &authRouter{ginRouter, authService}
}
