package alimail

import (
	"github.com/alibabacloud-go/tea/tea"
	"strings"
)

type MailMessage struct {
	queries map[string]interface{}
}

func NewMailMessage(accountName, subject string, toAddress ...string) *MailMessage {
	m := &MailMessage{queries: map[string]interface{}{}}
	m.AccountName(accountName).Subject(subject).ToAddress(toAddress...).ReplyToAddress(true).AddressType(1)
	return m
}

func (m *MailMessage) AccountName(accountName string) *MailMessage {
	m.queries["AccountName"] = tea.String(accountName)
	return m
}
func (m *MailMessage) AddressType(addressType int) *MailMessage {
	m.queries["AddressType"] = tea.Int(addressType)
	return m
}
func (m *MailMessage) ReplyToAddress(replyToAddress bool) *MailMessage {
	m.queries["ReplyToAddress"] = tea.Bool(replyToAddress)
	return m
}
func (m *MailMessage) ToAddress(toAddress ...string) *MailMessage {
	m.queries["ToAddress"] = tea.String(strings.Join(toAddress, ","))
	return m
}
func (m *MailMessage) Subject(subject string) *MailMessage {
	m.queries["Subject"] = tea.String(subject)
	return m
}
func (m *MailMessage) HtmlBody(htmlBody string) *MailMessage {
	m.queries["HtmlBody"] = tea.String(htmlBody)
	return m
}
func (m *MailMessage) TransFormToRequestParams() any {
	return m.queries
}
