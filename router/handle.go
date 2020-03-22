package router

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/iblockin/yunpian_err_repost/common"

	"github.com/iblockin/yunpian_err_repost/pkg/mysql"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// var Resu []Data

func Callbackhandle(c *gin.Context) {
	message := c.PostForm("sms_status")
	decodeMsg, err := url.QueryUnescape(message)
	if err != nil {
		logrus.Errorln("urlencode err", err)
	}
	decodeMsgProcess := "{\"data\": " + decodeMsg + "}"
	fmt.Println(decodeMsgProcess)
	var resp common.Resp
	err = json.Unmarshal([]byte(decodeMsgProcess), &resp)
	if err != nil {
		logrus.Errorln("Unmarshal data err:", err)
		return
	}
	newSlice := make([]common.Data, 0)
	for _, item := range resp.Result {
		if item.ReportStatus == "FAIL" {
			newSlice = append(newSlice, item)
		} else {
			continue
		}
	}
	logrus.Infoln("<test>", newSlice)
	for _, value := range newSlice {
		mysql.GetMysql().WriteData(value)
	}
}
