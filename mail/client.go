package mail

import (
	"errors"
	"github.com/pkg6/go-notify"
	"gopkg.in/gomail.v2"
)

type Client struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (c *Client) I() notify.IClient {
	return c
}

func (c *Client) Name() string {
	return notify.MailName
}

func (c *Client) Send(message notify.IMessage) (result notify.IResult) {
	result = notify.BuildResult(c.I(), message)
	params := message.TransFormToRequestParams()
	if gm, ok := params.(*gomail.Message); ok {
		d := gomail.NewDialer(c.Host, c.Port, c.Username, c.Password)
		if err := d.DialAndSend(gm); err != nil {
			result.WithException(err)
		}
	} else {
		result.WithException(errors.New("message.TransFormToRequestParams must be *gomail.Message"))
	}
	return result
}
