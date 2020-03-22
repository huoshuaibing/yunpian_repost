package main

import (
	"log"
	"net/http"
	"time"

	"github.com/iblockin/yunpian_err_repost/pkg/mysql"

	"github.com/iblockin/yunpian_err_repost/pkg/dingtalk"
	"github.com/iblockin/yunpian_err_repost/pkg/process"

	"github.com/iblockin/yunpian_err_repost/config"

	"os"

	"github.com/iblockin/yunpian_err_repost/router"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

func init() {
	file := "yunpian.log"
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Println(err)
	}
	logrus.SetOutput(f)
	logrus.SetLevel(logrus.InfoLevel)
	mysql.GetMysql().Init(config.GetConfig())
	ok := dingtalk.InstanceDingtalk.Init(config.GetConfig().Dingtalk)
	if ok != true {
		logrus.Errorf("dintalk init failure: %s", err)
		return
	}
}
func heartbeat() {
	heartbeaturl := config.GetConfig().Heartbeat
	client := http.Client{}
	client.Timeout = time.Second * time.Duration(120)
	for {
		_, err := client.Get(heartbeaturl)
		if err != nil {
			logrus.Errorln(err)
		}
		time.Sleep(time.Duration(20) * time.Second)
	}

}

func main() {
	go process.Send2dingtalk()
	go heartbeat()
	go dingtalk.InstanceDingtalk.Run()
	r := gin.Default()
	r.POST("/form_post", router.Callbackhandle)
	r.Run(":9091")
}
