package dingtalk

import "net/http"

type At struct {
	AtMobiles []string `json:"atmobiles"`
	IsAtAll   bool     `json:"isatall"`
}

type DTText struct {
	MsgType  string `json:"msgtype"`
	Markdown Text   `json:"markdown"`
	At       At     `json:"at"`
}

type Text struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type DingRobot struct {
	Url    string       `json:"url"`
	Client *http.Client `json:"client"`
	Data   chan Text    `json:"data"`
}
