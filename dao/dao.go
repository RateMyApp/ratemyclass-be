package dao

import (
	"context"

	"go.uber.org/fx"
)

type DbClient interface {
	Init() error
	Close(context.Context) error
}

var Module = fx.Module("dao",
	fx.Provide(
		fx.Annotate(
			NewPostgresClient,
			fx.As(new(DbClient)),
		)),
)
