/**
* Created by cks
* Date: 2020-12-02
* Time: 17:06
*/
package task

import (
	"SLU-System/config"
	"SLU-System/tools"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"

)

var RedisClient *redis.Client

func (task *Task) InitSubscribeRedisClient() (err error) {
	redisOpt := tools.RedisOption{
		Address:  config.Conf.Common.CommonRedis.RedisAddress,
		Password: config.Conf.Common.CommonRedis.RedisPassword,
		Db:       config.Conf.Common.CommonRedis.Db,
	}
	RedisClient = tools.GetRedisInstance(redisOpt)
	if pong, err := RedisClient.Ping().Result(); err != nil {
		logrus.Infof("RedisClient Ping Result pong: %s,  err: %s", pong, err)
	}

	go func() {
		redisSub := RedisClient.Subscribe(config.QueueName)
		ch := redisSub.Channel()
		for {
			msg, ok := <-ch
			if !ok {
				logrus.Debugf("redisSub Channel !ok: %v", ok)
				break
			}
			task.Push(msg.Payload)
		}
	}()
	return
}
