package main

import (
	"github.com/ratemyapp/config"
	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/routers"
	"github.com/ratemyapp/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(server.Module, config.Module, routers.Module, dao.Module).Run()
}
