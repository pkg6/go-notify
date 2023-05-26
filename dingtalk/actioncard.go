package dingtalk

import (
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

func (m ActionCardMessage) TransFormToRequestParams() any {
	m.Msgtype = "actionCard"
	return m
}

func NewActionCardMessage(title, text, btnOrientation, SingleTitle, singleURL string) notify.IMessage {
	m := ActionCardMessage{}
	m.ActionCard.Title = title
	m.ActionCard.Text = text
	m.ActionCard.BtnOrientation = btnOrientation
	m.ActionCard.SingleTitle = SingleTitle
	m.ActionCard.SingleURL = singleURL
	return m
}
