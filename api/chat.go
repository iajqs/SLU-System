/**
 * Created by cks
 * Date: 2020-11-29
 * Time: 17:37
 */
package api

type Chat struct {
}

func New() *Chat {
	return &Chat{}
}

// api server, Also, you can use gin, echo ... framework wrap
func (c *Chat) Run() {
	// init rpc client
	rpc.InitLogicRpcClient()

	r := router.Register()
	runMode := config.GetGinRunMode()
	logrus.Info("server start, now run mode is ", runMode)
	gin.SetMode(runMode)
	apiConfig := config.Conf.Api
	port := apiCOnfig.ApiBase.ListenPort
	flag.Parse()

	srv := &http.Server {
		Addr:	fmt.Sprintf(":%d", port),
		Hander: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("start listen : %s\n", err)
		}
	}()
	// if have two quit signal, this signal will priority capture, also can graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)
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