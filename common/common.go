package common

const (
	AlarmSend    int = 1
	AlarmNotSend int = 0
)

type Data struct {
	ErrorDetail     string `json:"error_detail"`
	Sid             int64  `json:"sid"`
	Uid             string `json:"uid"`
	UserReceiveTime string `json:"user_receive_time"`
	ErrorMsg        string `json:"error_msg"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
	AlarmStatus     int    `json:"alarmstatus"`
	CreateTime      string `json:"created_at"`
	UpdateTime      string `json:"updateed_at"`
}

type Resp struct {
	Result []Data `json:"data"`
}
