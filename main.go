package main

import (
	"github.com/ratemyapp/config"
	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/repositories"
	"github.com/ratemyapp/routers"
	"github.com/ratemyapp/server"
	"github.com/ratemyapp/services"
	"go.uber.org/fx"
)

func main() {
	fx.New(server.Module, config.Module, routers.Module, dao.Module, services.Module, repositories.Module).Run()
}
