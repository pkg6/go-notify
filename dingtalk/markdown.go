package dingtalk

import (
	"encoding/json"
	"github.com/pkg6/go-notify"
)

type MarkdownMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At At `json:"at"`
}

func (t MarkdownMessage) I() notify.IMessage {
	return &t
}

func (t *MarkdownMessage) AtMobiles(mobiles []string) {
	t.At.AtMobiles(mobiles)
}

func (t *MarkdownMessage) AtUserIds(userIds []string) {
	t.At.AtUserIds(userIds)
}

func (t *MarkdownMessage) AtAll() {
	t.At.AtAll()
}

func (t *MarkdownMessage) ToJson() string {
	t.Msgtype = "markdown"
	marshal, _ := json.Marshal(t)
	return string(marshal)
}
