package wework

import (
	"github.com/pkg6/go-notify"
)

type TextMessage struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
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
	m.Text.MentionedMobileList = mobiles
}

func (m *TextMessage) AtUserIds(userIds []string) {
	m.Text.MentionedList = append(m.Text.MentionedList, userIds...)
}

func (m *TextMessage) AtAll() {
	m.Text.MentionedList = append(m.Text.MentionedList, "@all")
}
