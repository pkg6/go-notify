package dingtalk

import (
	"encoding/json"
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

func (l LinkMessage) I() notify.IMessage {
	return &l
}

func (l *LinkMessage) AtMobiles(mobiles []string) {
	panic("No delivery required")
}

func (l *LinkMessage) AtUserIds(userIds []string) {
	panic("No delivery required")
}

func (l *LinkMessage) AtAll() {
	panic("No delivery required")
}

func (l *LinkMessage) ToJson() string {
	l.Msgtype = "link"
	marshal, _ := json.Marshal(l)
	return string(marshal)
}
