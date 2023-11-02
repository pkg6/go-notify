package wework

import (
	"errors"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-requests"
	"net/url"
	"strings"
)

// https://developer.work.weixin.qq.com/document/path/91770

const uri = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"

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
	return notify.NameWework
}

func (c *Client) url() *url.URL {
	var withKeyUrl *url.URL
	if strings.HasPrefix(c.Key, uri+"?key=") {
		withKeyUrl, _ = url.Parse(c.Key)
	} else {
		withKeyUrl, _ = url.Parse(uri)
		query := url.Values{}
		query.Set("key", c.Key)
		withKeyUrl.RawQuery = query.Encode()
	}
	return withKeyUrl
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
