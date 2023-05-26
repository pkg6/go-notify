package wework

import (
	"errors"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-requests"
	"net/url"
)

// https://developer.work.weixin.qq.com/document/path/91770

type Client struct {
	Key string
}
type WeWorkResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (c Client) I() notify.IClient {
	return &c
}

func (c Client) Name() string {
	return notify.WeworkName
}

func (c *Client) url() *url.URL {
	uri, _ := url.Parse("https://qyapi.weixin.qq.com/cgi-bin/webhook/send")
	query := url.Values{}
	query.Set("key", c.Key)
	uri.RawQuery = query.Encode()
	return uri
}
func (c *Client) Send(message notify.IMessage) (result notify.IResult) {
	var resp WeWorkResponse
	result = notify.BuildResult(c.I(), message)
	response, err := requests.PostJson(c.url().String(), message.TransFormToRequestParams())
	if err != nil {
		result.WithException(err)
		return result
	}
	err = response.Unmarshal(&resp)
	if err != nil {
		result.WithException(err)
		return result
	}
	if resp.Errcode != 0 {
		result.WithException(errors.New(resp.Errmsg))
		return result
	}
	result.WithResult(resp)
	return result
}
