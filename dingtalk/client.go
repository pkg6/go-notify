package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/gotool/client"
	"net/url"
	"strconv"
	"time"
)

//https://open.dingtalk.com/document/robots/custom-robot-access

type Client struct {
	AccessToken string
	Secret      string
}
type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (t Client) I() notify.IClient {
	return &t
}
func (t *Client) url() *url.URL {
	uri, _ := url.Parse("https://oapi.dingtalk.com/robot/send")
	query := url.Values{}
	query.Set("access_token", t.AccessToken)
	if t.Secret != "" {
		milliTimestamp := time.Now().UnixNano() / 1e6
		stringToSign := fmt.Sprintf("%s\n%s", strconv.Itoa(int(milliTimestamp)), t.Secret)
		mac := hmac.New(sha256.New, []byte(t.Secret))
		mac.Write([]byte(stringToSign))
		sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		query.Set("timestamp", strconv.Itoa(int(milliTimestamp)))
		query.Set("sign", sign)
	}
	uri.RawQuery = query.Encode()
	return uri
}
func (t *Client) Send(message notify.IMessage) (result notify.Result, err error) {
	var resp Response
	jsonByte, err := client.New(t.url().String()).PostJson(message.ToJson())
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(jsonByte, &resp)
	result = notify.BuildResult(t.I(), message, resp)
	return result, err
}
