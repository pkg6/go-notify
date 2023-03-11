package dingtalk

import (
	"encoding/json"
	"github.com/pkg6/go-notify"
)

type TextMessage struct {
	At   At `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Msgtype string `json:"msgtype"`
}

func (t TextMessage) I() notify.IMessage {
	return &t
}

func (t *TextMessage) AtMobiles(mobiles []string) {
	t.At.AtMobiles(mobiles)
}

func (t *TextMessage) AtUserIds(userIds []string) {
	t.At.AtUserIds(userIds)
}

func (t *TextMessage) AtAll() {
	t.At.AtAll()
}

func (t *TextMessage) ToJson() string {
	t.Msgtype = "text"
	marshal, _ := json.Marshal(t)
	return string(marshal)
}
