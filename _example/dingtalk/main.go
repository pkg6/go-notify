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
	message.Text.Content = "测试发送dingtalk"

	n := notify.New(client)
	sender := n.Send(message)
	for _, result := range sender {
		fmt.Println(fmt.Sprintf("%#v", result.Result()))
		fmt.Println(fmt.Sprintf("%#v", result.Status()))
		fmt.Println(fmt.Sprintf("%#v", result.Error()))
	}
}
