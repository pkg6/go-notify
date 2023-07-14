package alimail

import (
	"encoding/json"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pkg6/go-notify"
)

//https://help.aliyun.com/document_detail/29444.html?spm=a2c4g.29443.0.0.580c24cbh4QJBl

type Client struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	aliClient       *openapi.Client
}

func (c *Client) I() notify.IClient {
	var err error
	c.aliClient, err = openapi.NewClient(&openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: tea.String(c.AccessKeyId),
		// 必填，您的 AccessKey Secret
		AccessKeySecret: tea.String(c.AccessKeySecret),
		Endpoint:        tea.String("dm.aliyuncs.com"),
	})
	if err != nil {
		panic(err)
	}
	return c
}

func (c *Client) Name() string {
	return notify.NameAliMail
}

func (c *Client) commonParams() *openapi.Params {
	return &openapi.Params{
		// 接口名称
		Action: tea.String("SingleSendMail"),
		// 接口版本
		Version: tea.String("2015-11-23"),
		// 接口协议
		Protocol: tea.String("HTTPS"),
		// 接口 HTTP 方法
		Method:   tea.String("POST"),
		AuthType: tea.String("AK"),
		Style:    tea.String("RPC"),
		// 接口 PATH
		Pathname: tea.String("/"),
		// 接口请求体内容格式
		ReqBodyType: tea.String("json"),
		// 接口响应体内容格式
		BodyType: tea.String("json"),
	}
}

type Response struct {
	Body struct {
		EnvId     string `json:"EnvId"`
		RequestId string `json:"RequestId"`
	} `json:"body"`
	Headers struct {
		AccessControlAllowOrigin   string `json:"access-control-allow-origin"`
		AccessControlExposeHeaders string `json:"access-control-expose-headers"`
		Connection                 string `json:"connection"`
		ContentLength              string `json:"content-length"`
		ContentType                string `json:"content-type"`
		Date                       string `json:"date"`
		Etag                       string `json:"etag"`
		KeepAlive                  string `json:"keep-alive"`
		XAcsRequestId              string `json:"x-acs-request-id"`
		XAcsTraceId                string `json:"x-acs-trace-id"`
	} `json:"headers"`
	StatusCode int `json:"statusCode"`
}

func (c *Client) Send(message notify.IMessage) notify.IResult {
	result := notify.BuildResult(c.I(), message)
	params := message.TransFormToRequestParams()
	if queries, ok := params.(map[string]interface{}); ok {
		runtime := &util.RuntimeOptions{}
		request := &openapi.OpenApiRequest{
			Query: openapiutil.Query(queries),
		}
		api, err := c.aliClient.CallApi(c.commonParams(), request, runtime)
		if err != nil {
			result.WithException(err)
			return result
		}
		marshal, err := json.Marshal(api)
		if err != nil {
			result.WithException(err)
			return result
		}
		var response Response
		_ = json.Unmarshal(marshal, &response)
		result.WithResult(response)
	} else {
		result.WithException(errors.New("message.TransFormToRequestParams must be *gomail.Message"))
	}
	return result
}
