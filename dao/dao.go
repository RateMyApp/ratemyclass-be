package dao

import "go.uber.org/fx"

var Module = fx.Module("dao", fx.Provide(NewMongoDao))