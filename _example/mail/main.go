package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/gomail"
)

func main() {
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
