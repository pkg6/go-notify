package notify

// Sender 发送操作
func Sender(client IClient, message IMessage) (result Result, err error) {
	return client.Send(message)
}
