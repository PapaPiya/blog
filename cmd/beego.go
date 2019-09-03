package main

import (
	"beegolearn/pkg/common"
	"beegolearn/pkg/monitor"
	"beegolearn/pkg/service/app"
	"fmt"
	logger "go.uber.org/zap"
)

func main() {
    monitor.InitLog()
	defer func() {
		if err := recover(); err != nil {
		    fmt.Println(err)
			logger.L().Error("recover quit",logger.Error(err.(error)))
		}
	}()

	if err := common.InitConf(); err != nil {
        logger.L().Error("init conf error",logger.Error(err))
	}
	engine := app.Init()
	err := app.Run(engine)
	if err != nil {
		logger.L().Error("start service fail", logger.Error(err))
	}
	logger.L().Info("start service success")
/*	var g run.Group
	{
		// 初始化 gin 引擎
		engine := app.Init()
		// 启动服务
		g.Add(func() error {
			return app.Run(engine)
		},func(e error){
			logger.L().Error("gin stop", logger.Error(e))
		})
	}*/

/*	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}*/
}
