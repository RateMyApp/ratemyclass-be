package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ratemyapp/config"
	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/routers"
	"go.uber.org/fx"
)

func NewHttpServer(routers []routers.Router, lc fx.Lifecycle, appConfig config.AppConfig, ginRouter *gin.Engine, mongoDao *dao.MongoDao) *http.Server {

	for _, router := range routers {
		router.ExecRoutes()
	}

	srv := &http.Server{
		Addr:    ":" + appConfig.Port,
		Handler: ginRouter,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Println("Starting server on Port " + srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if  err := mongoDao.Client.Disconnect(ctx); err != nil {
				log.Fatal("Could not disconnect from mongo database successfully: " + err.Error())
			}

			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func NewGinRouter() *gin.Engine {
	return gin.Default()
}

var Module = fx.Module("server",
	fx.Provide(
		NewGinRouter,
		fx.Annotate(NewHttpServer, fx.ParamTags(`group:"routers"`)),
	),
	fx.Invoke(func(*http.Server) {}),
)
