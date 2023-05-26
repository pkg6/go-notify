package wework

import (
	"github.com/pkg6/go-notify"
)

type TemplateCardNewsNoticeMessage struct {
	Msgtype                string                 `json:"msgtype"`
	TemplateCardNewsNotice TemplateCardNewsNotice `json:"template_card"`
}

type TemplateCardNewsNotice struct {
	TemplateCardCommon
	CardImage struct {
		Url         string  `json:"url"`
		AspectRatio float64 `json:"aspect_ratio"`
	} `json:"card_image"`
	ImageTextArea struct {
		Type     int    `json:"type"`
		Url      string `json:"url"`
		Title    string `json:"title"`
		Desc     string `json:"desc"`
		ImageUrl string `json:"image_url"`
	} `json:"image_text_area"`
	VerticalContentList []struct {
		Title string `json:"title"`
		Desc  string `json:"desc"`
	} `json:"vertical_content_list"`
}

func NewTemplateCardNewsNoticeMessage(templateCard TemplateCardNewsNotice) notify.IMessage {
	m := TemplateCardNewsNoticeMessage{}
	m.TemplateCardNewsNotice = templateCard
	return m
}
func (m TemplateCardNewsNoticeMessage) TransFormToRequestParams() any {
	m.Msgtype = "template_card"
	return m
}
