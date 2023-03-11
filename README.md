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

~~~
package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/dingtalk"
)

func main() {
	client := dingtalk.DingtalkClient{AccessToken: ""}
	message := dingtalk.TextMessage{}
	message.Text.Content = "测试发送dingtalk"
	sender, err := notify.Sender(client.I(), message.I())
	fmt.Println(err)
	if response, ok := sender.Result.Response.(dingtalk.Response); ok {
		fmt.Println(response)
	}
}
~~~

## 自定义客户端

~~~
type IClient interface {
	I() IClient
	Send(message IMessage) (result Result, err error)
}
~~~

## 自定义消息内容

~~~
type IMessage interface {
	// I 初始化并返回IMessage
	I() IMessage
	// AtMobiles 消息需要@人员手机号
	AtMobiles(mobiles []string)
	// AtUserIds 消息需要@人员userid
	AtUserIds(userIds []string)
	// AtAll @所有人
	AtAll()
	// ToJson 转json字符串
	ToJson() string
}

~~~
