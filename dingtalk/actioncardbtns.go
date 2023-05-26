package dingtalk

import (
	"github.com/pkg6/go-notify"
)

type Btns struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}
type ActionCardBtnMessage struct {
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation string `json:"btnOrientation"`
		Btns           []Btns `json:"btns"`
	} `json:"actionCard"`
	Msgtype string `json:"msgtype"`
}

func NewActionCardBtnMessage(title, text, btnOrientation string, btns []Btns) notify.IMessage {
	m := &ActionCardBtnMessage{}
	m.ActionCard.Title = title
	m.ActionCard.Text = text
	m.ActionCard.BtnOrientation = btnOrientation
	m.ActionCard.Btns = btns
	return m
}
func (m ActionCardBtnMessage) TransFormToRequestParams() any {
	m.Msgtype = "actionCard"
	return m
}
