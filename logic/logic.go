/**
* Created by cks
* Date: 2020-12-01
* Time: 17:25
*/
package logic

import (
	"SLU-System/config"

	"github.com/sirupsen/logrus"
	
	"runtime"
)

type Logic struct {
}

func New() *Logic {
	return new(Logic)
}

func (logic *Logic) Run() {
	//read config
	logicConfig := config.Conf.Logic

	runtime.GOMAXPROCS(logicConfig.LogicBase.CpuNum)

	//init publish redis
	if err := logic.InitPublishRedisClient(); err != nil {
		logrus.Panicf("logic init publishRedisClient fail,err:%s", err.Error())
	}

	//init rpc server
	if err := logic.InitRpcServer(); err != nil {
		logrus.Panicf("logic init rpc server fail")
	}
}
