package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/dingtalk"
	"github.com/pkg6/go-notify/wework"
)

func main() {
	client := &wework.Client{Key: "693a91f6-7xxx-4bc4-97a0-0ec2sifa5aaa"}
	message := &dingtalk.TextMessage{}
	message.Text.Content = "测试发送wework"
	n := notify.New(client)
	sender := n.Send(message)
	for _, result := range sender {
		fmt.Println(fmt.Sprintf("%#v", result.Result()))
		fmt.Println(fmt.Sprintf("%#v", result.Status()))
		fmt.Println(fmt.Sprintf("%#v", result.Error()))
	}
}
