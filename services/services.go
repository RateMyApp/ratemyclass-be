package services

import "go.uber.org/fx"

var Module = fx.Module("services", fx.Provide(fx.Annotate(
	NewAuthServiceImpl, fx.As(new(AuthService)),
)))
