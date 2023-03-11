package dingtalk

import (
	"encoding/json"
	"github.com/pkg6/go-notify"
)

type FeedCardBtnMessage struct {
	Msgtype  string `json:"msgtype"`
	FeedCard struct {
		Links []struct {
			Title      string `json:"title"`
			MessageURL string `json:"messageURL"`
			PicURL     string `json:"picURL"`
		} `json:"links"`
	} `json:"feedCard"`
}

func (a FeedCardBtnMessage) I() notify.IMessage {
	return &a
}

func (a *FeedCardBtnMessage) AtMobiles(mobiles []string) {
	panic("No delivery required")
}

func (a *FeedCardBtnMessage) AtUserIds(userIds []string) {
	panic("No delivery required")
}

func (a *FeedCardBtnMessage) AtAll() {
	panic("No delivery required")
}

func (a *FeedCardBtnMessage) ToJson() string {
	a.Msgtype = "feedCard"
	marshal, _ := json.Marshal(a)
	return string(marshal)
}
