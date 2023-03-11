package notify

type Result struct {
	Client  IClient  `json:"client"`
	Message IMessage `json:"message"`
	Result  struct {
		Status    string `json:"status"`
		Response  any    `json:"response" `
		Exception error  `json:"exception"`
	} `json:"result"`
}

// BuildResult 在网关快速生成结构体
func BuildResult(client IClient, message IMessage, response any) Result {
	rt := Result{
		Client:  client,
		Message: message,
	}
	rt.Result.Response = response
	return rt
}
