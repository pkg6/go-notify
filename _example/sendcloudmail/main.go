package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/sendcloudmail"
)

func main() {
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
