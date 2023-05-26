package dingtalk

import (
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

func NewMarkdownMessage(title, text string) notify.IMessage {
	m := MarkdownMessage{}
	m.Markdown.Title = title
	m.Markdown.Text = text
	return m
}
func (m MarkdownMessage) TransFormToRequestParams() any {
	m.Msgtype = "markdown"
	return m
}

func (m *MarkdownMessage) AtMobiles(mobiles []string) {
	m.At.AtMobiles(mobiles)
}

func (m *MarkdownMessage) AtUserIds(userIds []string) {
	m.At.AtUserIds(userIds)
}

func (m *MarkdownMessage) AtAll() {
	m.At.AtAll()
}
