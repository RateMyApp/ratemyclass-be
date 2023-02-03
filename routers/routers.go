package routers

import "go.uber.org/fx"

type Router interface {
	ExecRoutes()
}

func AsRouter(f interface{}) any {
	return fx.Annotate(
		f,
		fx.As(new(Router)),
		fx.ResultTags(`group:"routers"`),
	)
}

var Module = fx.Module("routers",
	fx.Provide(AsRouter(NewHelloRouter)),
	fx.Provide(AsRouter(NewByeRouter)),
)
