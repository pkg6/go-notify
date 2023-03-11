package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/dingtalk"
)

func d() {
	client := dingtalk.Client{
		AccessToken: "45edd7aa6403682cbff302b4da591b648ac5dd0a88d768464ea2e3e9ac6e2089",
		Secret:      "SEC6abebb48b5547246af3f86d6c462eb2080af954ada7248d1ba5dcb0777b331f2",
	}
	message := dingtalk.TextMessage{}
	message.Text.Content = "测试发送dingtalk"
	sender, err := notify.Sender(client.I(), message.I())
	fmt.Println(err)
	if response, ok := sender.Result.Response.(dingtalk.Response); ok {
		fmt.Println(response)
	}
}

func main() {
	d()
}
