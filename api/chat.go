/**
 * Created by cks
 * Date: 2020-11-29
 * Time: 17:37
 */
package api

import (
	"SLU-System/api/rpc"
	"SLU-System/api/router"
	"SLU-System/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	
	"net/http"
	"context"
	"time"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Chat struct {
}

func New() *Chat {
	return &Chat{}
}

// api server, Also, you can use gin, echo ... framework wrap
func (c *Chat) Run() {
	fmt.Println("initing... the logic rpc client")
	// init rpc client
	rpc.InitLogicRpcClient()
	fmt.Println("inited the logic rpc client")
	r := router.Register()
	runMode := config.GetGinRunMode()
	logrus.Info("server start, now run mode is ", runMode)
	gin.SetMode(runMode)
	apiConfig := config.Conf.Api
	port := apiConfig.ApiBase.ListenPort
	bind := apiConfig.ApiBase.Bind
	flag.Parse()

	srv := &http.Server {
		Addr:	 fmt.Sprintf(":%d", port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("start listen : %s\n", err)
		}
	}()
	// if have two quit signal, this signal will priority capture, also can graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	logrus.Infof("Shutdown Server ...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("Server Shutdown:", err)
	}
	logrus.Infof("Server exiting")
	os.Exit(0)
}