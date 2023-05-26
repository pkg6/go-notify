package wework

import (
	"github.com/pkg6/go-notify"
)

type NewsMessage struct {
	Msgtype string `json:"msgtype"`
	News    struct {
		Articles []Article `json:"articles"`
	} `json:"news"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
}

func NewNewsMessage(articles []Article) notify.IMessage {
	m := NewsMessage{}
	m.News.Articles = articles
	return m
}
func (m NewsMessage) TransFormToRequestParams() any {
	m.Msgtype = "news"
	return m
}
