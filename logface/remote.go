package logface

import (
	"fmt"
	"github.com/b1b2b3b4b5b6/goc/tl/jsont"
	"github.com/b1b2b3b4b5b6/goc/tl/errt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type remoteWirte struct {
}

func (p remoteWirte) Write(data []byte) (n int, err error) {
	chRemote <- string(data)
	return os.Stdout.Write(data)
}

var rw remoteWirte

var chRemote = make(chan string, 1000)

func init() {
	go remoteLoop()
}

func remoteLoop() {
	json, err := cfg.TakeJson("LogRedisCfg")
	if err != nil {
		logrus.Warn("no cfg, remote log not init")
		for {
			<-chRemote
		}
	}
	var remoteCfg struct {
		Host string
	}
	err = jsont.Decode(json, &remoteCfg)
	errt.Errpanic(err)

	pool := redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", remoteCfg.Host) },
	}

	for {
		logStr := <-chRemote
		rep, err := pool.Get().Do("LPUSH", "log", logStr)
		if err != nil {
			logrus.Warn(fmt.Sprintf("send remote log fail[%+v, %+v]", rep, err))
		}
	}
}
