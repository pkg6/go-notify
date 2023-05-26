package wework

import (
	"github.com/pkg6/go-notify"
)

type MarkdownMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

func NewMarkdownMessage(content string) notify.IMessage {
	m := MarkdownMessage{}
	m.Markdown.Content = content
	return m
}
func (m MarkdownMessage) TransFormToRequestParams() any {
	m.Msgtype = "markdown"
	return m
}
