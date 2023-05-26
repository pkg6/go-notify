package dingtalk

import (
	"github.com/pkg6/go-notify"
)

type LinkMessage struct {
	Msgtype string `json:"msgtype"`
	Link    struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicURL     string `json:"picUrl"`
		MessageURL string `json:"messageUrl"`
	} `json:"link"`
}

func NewLinkMessage(text, title, picURL, messageURL string) notify.IMessage {
	m := LinkMessage{}
	m.Link.Text = text
	m.Link.Title = title
	m.Link.PicURL = picURL
	m.Link.MessageURL = messageURL
	return m
}

func (m LinkMessage) TransFormToRequestParams() any {
	m.Msgtype = "link"
	return m
}
