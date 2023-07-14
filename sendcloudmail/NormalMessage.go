package sendcloudmail

import (
	"strings"
)

type NormalMessage struct {
	Params map[string]string
}

func NewNormalMessage(form, subject string, to ...string) *NormalMessage {
	m := &NormalMessage{Params: map[string]string{}}
	m.SetForm(form).SetSubject(subject).SetTo(to...)
	return m
}

func (m *NormalMessage) SetForm(form string) *NormalMessage {
	m.Params["from"] = form
	return m
}
func (m *NormalMessage) SetTo(tos ...string) *NormalMessage {
	m.Params["to"] = strings.Join(tos, ";")
	return m
}
func (m *NormalMessage) SetCc(cc ...string) *NormalMessage {
	m.Params["cc"] = strings.Join(cc, ";")
	return m
}

func (m *NormalMessage) SetBcc(bcc ...string) *NormalMessage {
	m.Params["bcc"] = strings.Join(bcc, ";")
	return m
}
func (m *NormalMessage) SetSubject(subject string) *NormalMessage {
	m.Params["subject"] = subject
	return m
}
func (m *NormalMessage) Html(html string) *NormalMessage {
	m.Params["html"] = html
	return m
}
func (m *NormalMessage) ContentSummary(contentSummary string) *NormalMessage {
	m.Params["contentSummary"] = contentSummary
	return m
}
func (m *NormalMessage) FromName(fromName string) *NormalMessage {
	m.Params["fromName"] = fromName
	return m
}
func (m *NormalMessage) TransFormToRequestParams() any {
	return m.Params
}
