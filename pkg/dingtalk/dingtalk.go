package dingtalk

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// fix add mutex
var T Text

var InstanceDingtalk DingRobot

func (d *DingRobot) Init(Url string) bool {
	if len(Url) == 0 {
		logrus.Println("url length is 0")
		return false
	}
	d.Url = Url
	d.Data = make(chan Text, 100)
	d.Client = new(http.Client)
	return true
}

func (d *DingRobot) Run() {
	for newmsg := range d.Data {
		msg := DTText{}
		msg.MsgType = "markdown"
		msg.Markdown.Title = newmsg.Title
		msg.Markdown.Text = newmsg.Text
		msg.At.AtMobiles = []string{}
		msg.At.IsAtAll = false
		jsonByte, err := json.Marshal(msg)
		if err != nil {
			logrus.Errorln("marshal json err", err)
			continue
		}
		jsonString := string(jsonByte)
		logrus.Infof("%s", jsonString)
		req, err := http.NewRequest("POST", d.Url, strings.NewReader(jsonString))
		if err != nil {
			logrus.Errorln("new request err", err)
			continue
		}
		req.Header.Add("Accept-Charset", "utf-8")
		req.Header.Add("Content-Type", "application/json")
		_, err = d.Client.Do(req)
		d.Client.Timeout = time.Duration(30) * time.Second
		if err != nil {
			logrus.Errorln(err)
		}
	}
}
