package dingtalk

import (
	"github.com/pkg6/go-notify"
)

type TextMessage struct {
	At   At `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Msgtype string `json:"msgtype"`
}

func NewTextMessage(content string) notify.IMessage {
	m := TextMessage{}
	m.Text.Content = content
	return m
}
func (m TextMessage) TransFormToRequestParams() any {
	m.Msgtype = "text"
	return m
}

func (m *TextMessage) AtMobiles(mobiles []string) {
	m.At.AtMobiles(mobiles)
}

func (m *TextMessage) AtUserIds(userIds []string) {
	m.At.AtUserIds(userIds)
}

func (m *TextMessage) AtAll() {
	m.At.AtAll()
}
