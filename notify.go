package notify

const (
	NameDingTalk          = "dingtalk"
	NameWework            = "wework"
	NameGoMail            = "gomail"
	NameSendCloudMail     = "sendcloudmail"
	NameAliMail           = "alimail"
	RESULT_STATUS_ERROR   = 1
	RESULT_STATUS_SUCCESS = 0
)

func New(clients ...IClient) *Notify {
	n := &Notify{
		clients: map[string]IClient{},
		names:   []string{},
	}
	for _, client := range clients {
		n.Extend(client)
	}
	return n
}

// Notify 通知实体
type Notify struct {
	names   []string
	clients map[string]IClient
}

// Extend 注册自定义网关
func (n *Notify) Extend(client IClient, names ...string) *Notify {
	name := client.Name()
	if len(names) > 0 {
		name = names[0]
	}
	n.names = append(n.names, name)
	n.clients[name] = client
	return &Notify{names: names, clients: n.clients}
}

// Names 使用别名发送
func (n *Notify) Names(names ...string) *Notify {
	return &Notify{names: names, clients: n.clients}
}

// Send 发送
func (n *Notify) Send(message IMessage) map[string]IResult {
	results := make(map[string]IResult)
	for _, name := range n.names {
		if client, ok := n.clients[name]; ok {
			results[name] = client.Send(message)
		}
	}
	return results
}

// SendFirstName 发送
func (n *Notify) SendFirstName(message IMessage) IResult {
	name := n.names[0]
	client, ok := n.clients[name]
	if !ok {
		return nil
	}
	return client.Send(message)
}
