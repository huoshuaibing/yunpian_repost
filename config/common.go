package config

type Database struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Protocol  string `json:"protocol"`
	Ip        string `json:"ip"`
	Prot      int    `json:"port"`
	Basename  string `json:"basename"`
	DataTable string `json:"datatable"`
	Param     string `json:"param"`
	Value     string `json:"value"`
	Key       string `json:"key"`
}

type Config struct {
	Dingtalk  string   `json:"dingtalk"`
	Heartbeat string   `json:"heartbeat"`
	Database  Database `json:"database"`
}
