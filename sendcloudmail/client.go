package sendcloudmail

import (
	"errors"
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-requests"
	"net/url"
)

var SendUrl = "https://api.sendcloud.net/apiv2/mail/send"

type Client struct {
	ApiUser string `json:"api_user"`
	ApiKey  string `json:"api_key"`
}

func (c *Client) I() notify.IClient {
	return c
}

func (c *Client) Name() string {
	return notify.NameSendCloudMail
}

type Response struct {
	StatusCode int `json:"statusCode"`
	Info       struct {
		EmailIdList []string `json:"emailIdList"`
	} `json:"info"`
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

func (c *Client) Send(message notify.IMessage) (result notify.IResult) {
	result = notify.BuildResult(c.I(), message)
	params := message.TransFormToRequestParams()

	if gm, ok := params.(map[string]string); ok {
		value := url.Values{}
		value.Set("apiUser", c.ApiUser)
		value.Set("apiKey", c.ApiKey)
		for s, s2 := range gm {
			value.Set(s, s2)
		}
		form, err := requests.PostForm(SendUrl, value)
		if err != nil {
			result.WithException(err)
			return nil
		}
		var response Response
		if err = form.Unmarshal(&response); err != nil {
			result.WithException(err)
			return nil
		}
		if response.StatusCode != 200 {
			result.WithException(fmt.Errorf(response.Message))
			return result
		}
		result.WithResult(response)
	} else {
		result.WithException(errors.New("message.TransFormToRequestParams must be *gomail.Message"))
	}
	return result
}
