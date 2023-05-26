<h1 align="center">Go Notify</h1>

<p align="center">:calling: 一款满足你的多种渠道消息通知</p>


[![Go Report Card](https://goreportcard.com/badge/github.com/pkg6/go-notify)](https://goreportcard.com/report/github.com/pkg6/go-notify)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/pkg6/go-notify?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/pkg6/go-notify/-/badge.svg)](https://sourcegraph.com/github.com/pkg6/go-notify?badge)
[![Release](https://img.shields.io/github/release/pkg6/go-notify.svg?style=flat-square)](https://github.com/pkg6/go-notify/releases)


## 安装

```
$ go get github.com/pkg6/go-notify
```

## 使用



## 平台支持

* [钉钉群机器人](https://developers.dingtalk.com/document/app/custom-robot-access)
* [微信群机器人](https://developer.work.weixin.qq.com/document/path/91770)
* [邮件](https://github.com/go-gomail/gomail)



## 使用
<details>
<summary><b>钉钉群机器人</b></summary>

```
package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/dingtalk"
)

func main() {
	client := &dingtalk.Client{
		AccessToken: "27bbe68cc8b57acc2973b59fd7ae2460fb0b2322ce2e8660f5fb5b75aee04e88",
		Secret:      "SEC55f77c19089ef4aee0be143a77d12730f2daaa2390b212cffb1e1ac1f23f8ccc",
	}
	message := &dingtalk.TextMessage{}
	message.Text.Content = "测试发送dingtalk2"
	n := notify.New(client)
	sender := n.Send(message)
	for _, result := range sender {
		fmt.Println(fmt.Sprintf("%#v", result.Result()))
		fmt.Println(fmt.Sprintf("%#v", result.Status()))
		fmt.Println(fmt.Sprintf("%#v", result.Error()))
	}
}
```
</details>

<details>
<summary><b>微信群机器人</b></summary>

```
package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/wework"
)

func main() {
	client := &wework.Client{
		Key: "693a91f6-7xxx-4bc4-97a0-0ec2sifa5aaa",
	}
	message := &wework.TextMessage{}
	message.Text.Content = "测试发送wework"
	n := notify.New(client)
	sender := n.Send(message)
	for _, result := range sender {
		fmt.Println(fmt.Sprintf("%#v", result.Result()))
		fmt.Println(fmt.Sprintf("%#v", result.Status()))
		fmt.Println(fmt.Sprintf("%#v", result.Error()))
	}
}
```
</details>



<details>
<summary><b>邮件</b></summary>

```
package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/mail"
)

func main() {
	client := &mail.Client{
		Host:     "smtp.126.com",
		Port:     456,
		Username: "******@126.com",
		Password: "***************",
	}
	message := &mail.Message{}
	message.SetForm("********@126.com")
	message.SetTo("**********@qq.com")
	message.Html("<h3>GO-Notify</h3><p>欢迎使用GO-Notify</p>")
	n := notify.New(client)
	sender := n.Send(message)
	for _, result := range sender {
		fmt.Println(fmt.Sprintf("%#v", result.Result()))
		fmt.Println(fmt.Sprintf("%#v", result.Status()))
		fmt.Println(fmt.Sprintf("%#v", result.Error()))
	}
}
```
</details>


## 自定义客户端

~~~
type IClient interface {
	I() IClient
	Name() string
	Send(message IMessage) IResult
}
~~~

## 自定义消息内容

~~~
// IMessage 消息
type IMessage interface {
	TransFormToRequestParams() any
}
~~~
