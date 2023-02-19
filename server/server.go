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

func NewHttpServer(routers []routers.Router, lc fx.Lifecycle, appConfig config.AppConfig, ginRouter *gin.Engine, dbClient dao.DbClient) *http.Server {
	if err := dbClient.Init(); err != nil {
		log.Panic("Error Connecting to Database: ", err)
	}
	for _, router := range routers {
		router.ExecRoutes()
	}

	srv := &http.Server{
		Addr:    ":" + appConfig.PORT,
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
			if err := dbClient.Close(ctx); err != nil {
				log.Fatal("Could not disconnect from database successfully: " + err.Error())
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
