package dingtalk

import (
	"encoding/json"
	"github.com/pkg6/go-notify"
)

type ActionCardMessage struct {
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation string `json:"btnOrientation"`
		SingleTitle    string `json:"singleTitle"`
		SingleURL      string `json:"singleURL"`
	} `json:"actionCard"`
	Msgtype string `json:"msgtype"`
}

func (a ActionCardMessage) I() notify.IMessage {
	return &a
}

func (a *ActionCardMessage) AtMobiles(mobiles []string) {
	panic("No delivery required")
}

func (a *ActionCardMessage) AtUserIds(userIds []string) {
	panic("No delivery required")
}

func (a *ActionCardMessage) AtAll() {
	panic("No delivery required")
}

func (a *ActionCardMessage) ToJson() string {
	a.Msgtype = "actionCard"
	marshal, _ := json.Marshal(a)
	return string(marshal)
}
