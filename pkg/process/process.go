package process

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/iblockin/yunpian_err_repost/pkg/dingtalk"
	"github.com/iblockin/yunpian_err_repost/pkg/mysql"
)

func Send2dingtalk() {
	for {
		dataArray := mysql.GetMysql().GetData()
		for _, x := range dataArray {
			if x.ErrorDetail == "防骚扰号码" || x.ErrorDetail == "运营商内容驳回" || x.ErrorDetail == "来自MSC的未知错误。" || x.ErrorDetail == "超过月最大发送MT数量" {
				dingtalk.T.Title = "用户短信接收失败"
				content := "# " + dingtalk.T.Title + "\n\n" +
					"------------------------------\n\n" +
					"### 接收状态: \n\n" + x.ReportStatus + "\n\n" +
					"### 用户接收时间: \n\n" + x.UserReceiveTime + "\n\n" +
					"### 用户手机号: \n\n" + x.Mobile + "\n\n" +
					"### 用户id: \n\n" + x.Uid + "\n\n" +
					"### 失败原因:\n\n" + x.ErrorDetail + "\n\n" +
					"------------------------------\n\n"
				dingtalk.T.Text = content
				dingtalk.InstanceDingtalk.Data <- dingtalk.T
				status := mysql.GetMysql().UpdateData(x.Sid)
				if !status {
					logrus.Errorln("update data err")
					continue
				}
			}
		}
		time.Sleep(time.Duration(5) * time.Second)
	}
}
