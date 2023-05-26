package wework

import (
	"github.com/pkg6/go-notify"
)

type TemplateCardTextNoticeMessage struct {
	Msgtype                string                 `json:"msgtype"`
	TemplateCardTextNotice TemplateCardTextNotice `json:"template_card"`
}

type TemplateCardTextNotice struct {
	TemplateCardCommon
	EmphasisContent struct {
		Title string `json:"title"`
		Desc  string `json:"desc"`
	} `json:"emphasis_content"`
	SubTitleText string `json:"sub_title_text"`
}

func NewTemplateCardTextNoticeMessage(templateCard TemplateCardTextNotice) notify.IMessage {
	m := TemplateCardTextNoticeMessage{}
	m.TemplateCardTextNotice = templateCard
	return m
}

func (m TemplateCardTextNoticeMessage) TransFormToRequestParams() any {
	m.Msgtype = "template_card"
	return m
}
