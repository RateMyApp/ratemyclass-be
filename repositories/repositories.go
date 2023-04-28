package repositories

import "go.uber.org/fx"

var Module = fx.Module("repositories", fx.Provide(fx.Annotate(
	NewUserRepository, fx.As(new(UserRepository)),
)))