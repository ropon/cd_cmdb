package routers

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/ropon/cd_cmdb/conf"
	"github.com/ropon/cd_cmdb/controllers"
	_ "github.com/ropon/cd_cmdb/docs"
	"github.com/ropon/cd_cmdb/utils"
	"github.com/ropon/logger"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// setupRouter 初始化路由
func setupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(ginLogger())
	engine.Use(cors())
	engine.Use(utils.TraceHttpRoot(conf.SERVERNAME, conf.Cfg.External["JaegerAgentAddr"]))
	engine.Use(gin.Recovery())
	pprof.Register(engine)

	engine.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	backend := engine.Group("/backend")
	{
		backend.Any(":server/*action", controllers.HttpProxy)
	}
	v1 := engine.Group("/cd_cmdb/api/v1")
	{
		v1.GET("/hi", controllers.Hi)

		v1.POST("/service", controllers.CreateService)
		v1.DELETE("/service/:id", controllers.DeleteService)
		v1.PUT("/service/:id", controllers.UpdateService)
		v1.PATCH("/service/:id", controllers.PatchUpdateService)
		v1.GET("/service", controllers.GetServices)
		v1.GET("/service/:id", controllers.GetService)

		v1.GET("/server", controllers.GetServers)
		v1.GET("/server/:id", controllers.GetServer)

		v1.GET("/db", controllers.GetDbs)
		v1.GET("/db/:id", controllers.GetDb)

		v1.GET("/test", controllers.Test)
	}

	engine.NoRoute(func(c *gin.Context) {
		utils.GinErrRsp(c, utils.ErrCodeGeneralFail, "Page not found")
	})
	return engine
}

func Run(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: setupRouter(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("listen error: %s\n", err.Error())
		}
	}()
	fmt.Printf("[GIN-debug] Listening and serving HTTP on %s\n", addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	<-quit
	logger.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server Shutdown error: %s", err.Error())
	}
	fmt.Printf("\nServer exiting\n")
}
