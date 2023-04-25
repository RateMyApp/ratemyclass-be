package routers

import "go.uber.org/fx"
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
)
