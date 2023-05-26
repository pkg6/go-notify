package dingtalk

import (
	"github.com/pkg6/go-notify"
)

type Link struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}
type FeedCardBtnMessage struct {
	Msgtype  string `json:"msgtype"`
	FeedCard struct {
		Links []Link `json:"links"`
	} `json:"feedCard"`
}

func NewFeedCardBtnMessage(link []Link) notify.IMessage {
	m := FeedCardBtnMessage{}
	m.FeedCard.Links = link
	return m
}

func (m *FeedCardBtnMessage) AddLink(link Link) {
	m.FeedCard.Links = append(m.FeedCard.Links, link)
}

func (m FeedCardBtnMessage) TransFormToRequestParams() any {
	m.Msgtype = "feedCard"
	return m
}
