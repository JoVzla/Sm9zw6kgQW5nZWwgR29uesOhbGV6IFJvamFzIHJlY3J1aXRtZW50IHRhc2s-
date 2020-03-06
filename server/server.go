package server

import (
	"context"
	"gogoapps/logger"
	"gogoapps/middlewares"
	"gogoapps/models/config"
	"gogoapps/routes"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var srv *http.Server

func Start(cach *cache.Cache, configuration config.Configuration) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(cors.Default())
	r.Use(middlewares.InitConnections(cach, configuration))

	routes.Weather(r)

	srv = &http.Server{
		Addr:    ":" + configuration.Port,
		Handler: r,
	}

	logger.Log.Info("Starting rest server")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Log.Error("Error starting rest server " + err.Error())
	}

}

//Stop the rest server
func Stop() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 3 seconds.
	select {
	case <-ctx.Done():
		logger.Log.Info("timeout of 3 seconds.")
	}
	logger.Log.Info("Server exiting")
}
