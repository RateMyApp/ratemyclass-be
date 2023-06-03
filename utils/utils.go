package utils

import "go.uber.org/fx"

var Module = fx.Module("utils", fx.Provide(
	fx.Annotate(
		NewJwtUtil,
		fx.As(new(JwtUtil)),
	),
	fx.Annotate(
		NewTimeUtil,
		fx.As(new(TimeUtil)),
	),
),
)
