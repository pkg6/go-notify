package main

import (
	"fmt"
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/alimail"
)

func main() {
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
