package dingtalk

import (
	"encoding/json"
	"github.com/pkg6/go-notify"
)

type ActionCardBtnMessage struct {
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation string `json:"btnOrientation"`
		Btns           []struct {
			Title     string `json:"title"`
			ActionURL string `json:"actionURL"`
		} `json:"btns"`
	} `json:"actionCard"`
	Msgtype string `json:"msgtype"`
}

func (a ActionCardBtnMessage) I() notify.IMessage {
	return &a
}

func (a *ActionCardBtnMessage) AtMobiles(mobiles []string) {
	panic("No delivery required")
}

func (a *ActionCardBtnMessage) AtUserIds(userIds []string) {
	panic("No delivery required")
}

func (a *ActionCardBtnMessage) AtAll() {
	panic("No delivery required")
}

func (a *ActionCardBtnMessage) ToJson() string {
	a.Msgtype = "actionCard"
	marshal, _ := json.Marshal(a)
	return string(marshal)
}
