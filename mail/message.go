package mail

import (
	"github.com/pkg6/go-notify"
	"gopkg.in/gomail.v2"
)

type Message struct {
	Form            string
	To              []string
	Cc              string
	Bcc             string
	Subject         string
	BodyContentType string
	Body            string
	Attach          []string
}

func (m *Message) SetForm(form string) notify.IMessage {
	m.Form = form
	return m
}
func (m *Message) SetTo(tos ...string) notify.IMessage {
	for _, to := range tos {
		m.To = append(m.To, to)
	}
	return m
}
func (m *Message) SetCc(cc string) notify.IMessage {
	m.Cc = cc
	return m
}

func (m *Message) SetBcc(bcc string) notify.IMessage {
	m.Bcc = bcc
	return m
}
func (m *Message) SetSubject(subject string) notify.IMessage {
	m.Subject = subject
	return m
}
func (m *Message) Html(html string) notify.IMessage {
	m.BodyContentType = "text/html"
	m.Body = html
	return m
}
func (m *Message) Text(text string) notify.IMessage {
	m.BodyContentType = "text/plain"
	m.Body = text
	return m
}

func (m *Message) TransFormToRequestParams() any {
	message := gomail.NewMessage()
	message.SetHeader("From", m.Form)
	message.SetHeader("To", m.To...)
	message.SetHeader("Subject", m.Subject)
	message.SetHeader("Cc", m.Cc)
	message.SetHeader("Bcc", m.Bcc)
	message.SetBody(m.BodyContentType, m.Body)
	for _, attach := range m.Attach {
		message.Attach(attach)
	}
	return message
}
