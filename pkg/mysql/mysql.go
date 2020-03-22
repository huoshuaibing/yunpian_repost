package mysql

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/iblockin/yunpian_err_repost/common"
	"github.com/sirupsen/logrus"

	"github.com/iblockin/yunpian_err_repost/config"
)

var onceMysql sync.Once
var insMysql *Mysql

type Mysql struct {
	SchemaYunpian *sql.DB
	lockYunpian   sync.Mutex
}

const (
	AlarmSend    int = 1
	AlarmNotSend int = 0
)

func GetMysql() *Mysql {
	onceMysql.Do(func() {
		insMysql = new(Mysql)
	})
	return insMysql
}

const PayOut string = "2006-01-02 15:04:05"

func (m *Mysql) Init(cfg *config.Config) {
	OpenSrcDbCommand := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s=%s",
		cfg.Database.Username, cfg.Database.Password,
		cfg.Database.Protocol, cfg.Database.Ip,
		int(cfg.Database.Prot), cfg.Database.Basename,
		cfg.Database.Param, cfg.Database.Value)
	var err error
	m.SchemaYunpian, err = sql.Open("mysql", OpenSrcDbCommand)
	if err != nil {
		logrus.Errorln("Open Src Error.", err)
		return
	}
}
func (m *Mysql) WriteData(d common.Data) {
	m.lockYunpian.Lock()
	defer m.lockYunpian.Unlock()
	InsertCommand := fmt.Sprintf(`insert into %s (errordetail,sid,uid,userreceivetime,errormsg,mobile,reportstatus,alarmstatus,created_at,updated_at) values ('%v','%v','%v','%v','%v','%v','%v','%v','%v','%v')`,
		config.ConfigInstance.Database.DataTable, d.ErrorDetail, d.Sid, d.Uid, d.UserReceiveTime, d.ErrorMsg, d.Mobile, d.ReportStatus, AlarmNotSend, time.Now().Format(PayOut), time.Now().Format(PayOut))
	fmt.Println(InsertCommand)
	_, err := m.SchemaYunpian.Exec(InsertCommand)
	if err != nil {
		logrus.Errorln("Insert mysql err.", err)
		return
	}
}

// func (m *Mysql) QueryData(sid int64) bool {
// 	m.lockYunpian.Lock()
// 	defer m.lockYunpian.Unlock()
// 	SelectCommand := fmt.Sprintf(`select 1 from %v where sid = %d limit 1;`, config.ConfigInstance.Database.DataTable, sid)
// 	rows := m.SchemaYunpian.QueryRow(SelectCommand)
// 	if rows == nil {
// 		logrus.Infoln("sou bu dao")
// 		return false
// 	}
// 	logrus.Infoln("neng sou  dao")
// 	logrus.Infoln(rows)
// 	return true
// }

func (m *Mysql) GetData() []common.Data {
	m.lockYunpian.Lock()
	defer m.lockYunpian.Unlock()
	tmpSlice := make([]common.Data, 0)
	GetCommand := fmt.Sprintf(`select errordetail,sid,uid,userreceivetime,errormsg,mobile,reportstatus,alarmstatus,created_at,updated_at from %v where AlarmStatus = 0;`, config.ConfigInstance.Database.DataTable)
	rows, err := m.SchemaYunpian.Query(GetCommand)
	closeRows := func() {
		if rows != nil {
			rows.Close()
		}
	}
	defer closeRows()
	if err != nil || rows == nil {
		return tmpSlice
	}
	var sum int = 0
	var ErrorDetail string
	var Sid int64
	var Uid string
	var UserReceiveTime string
	var ErrorMsg string
	var Mobile string
	var ReportStatus string
	var AlarmStatus int
	var CreateTime string
	var UpdateTime string
	for rows.Next() {
		sum++
		err := rows.Scan(&ErrorDetail, &Sid, &Uid, &UserReceiveTime, &ErrorMsg, &Mobile, &ReportStatus, &AlarmStatus, &CreateTime, &UpdateTime)
		if err != nil {
			logrus.Errorln("Scan rows error", err)
			return tmpSlice
		}
		data := common.Data{
			ErrorDetail:     ErrorDetail,
			Sid:             Sid,
			Uid:             Uid,
			UserReceiveTime: UserReceiveTime,
			ErrorMsg:        ErrorMsg,
			Mobile:          Mobile,
			ReportStatus:    ReportStatus,
			AlarmStatus:     AlarmStatus,
			CreateTime:      CreateTime,
			UpdateTime:      UpdateTime,
		}
		tmpSlice = append(tmpSlice, data)
	}
	return tmpSlice
}
func (m *Mysql) UpdateData(id int64) bool {
	logrus.Infoln("Id is :", id)
	m.lockYunpian.Lock()
	defer m.lockYunpian.Unlock()
	UpdateCommand := fmt.Sprintf(`update %s set alarmstatus=%d, updated_at='%v' where sid= '%v';`, config.ConfigInstance.Database.DataTable, common.AlarmSend, time.Now().Format(PayOut), id)
	_, err := m.SchemaYunpian.Exec(UpdateCommand)
	if err != nil {
		logrus.Errorln("UPdate table err", err)
		return false
	}
	return true
}