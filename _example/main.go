package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/alimail"
	"github.com/pkg6/go-notify/dingtalk"
	"github.com/pkg6/go-notify/gomail"
	"github.com/pkg6/go-notify/sendcloudmail"
	"github.com/pkg6/go-notify/wework"
)

func dingtalkSend() {
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

func weworkSend() {
	client := &wework.Client{
		Key: "693a91f6-7xxx-4bc4-97a0-0ec2sifa5aaa",
	}
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
func mailSend() {
	client := &gomail.Client{
		Host:     "smtp.126.com",
		Port:     456,
		Username: "******@126.com",
		Password: "***************",
	}
	message := &gomail.Message{}
	message.SetForm("*********@126.com")
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

func sendCloudMail() {
	client := &sendcloudmail.Client{ApiKey: "", ApiUser: ""}
	message := sendcloudmail.NewNormalMessage("form-pkg6@github.com", "Go Notify", "to-pkg6@github.com")
	message.Html("Go Notify")
	n := notify.New(client)
	sender := n.Send(message)
	for _, result := range sender {
		fmt.Println(fmt.Sprintf("%#v", result.Result()))
		fmt.Println(fmt.Sprintf("%#v", result.Status()))
		fmt.Println(fmt.Sprintf("%#v", result.Error()))
	}

}

func alimailSend() {
	client := &alimail.Client{AccessKeyId: "", AccessKeySecret: ""}
	message := alimail.NewMailMessage("form-pkg6@github.com", "Go Notify", "to-pkg6@github.com")
	message.HtmlBody("Go Notify")
	n := notify.New(client)
	sender := n.Send(message)
	for _, result := range sender {
		fmt.Println(fmt.Sprintf("%#v", result.Result()))
		fmt.Println(fmt.Sprintf("%#v", result.Status()))
		fmt.Println(fmt.Sprintf("%#v", result.Error()))
	}

}
func main() {
	//mailSend()
	alimailSend()
}
