package notify

const (
	DingTalkName          = "dingtalk"
	WeworkName            = "wework"
	RESULT_STATUS_ERROR   = 1
	RESULT_STATUS_SUCCESS = 0
)

func New(clients ...IClient) *Notify {
	n := new(Notify).clone()
	for _, client := range clients {
		n.Extend(client.Name(), client)
	}
	return n
}

// Notify 通知实体
type Notify struct {
	names   []string
	clients map[string]IClient
}

// 初始化参数
func (n Notify) clone() *Notify {
	n.clients = make(map[string]IClient)
	return &n
}

// Extend 注册自定义网关
func (n *Notify) Extend(name string, client IClient) {
	n.names = append(n.names, name)
	n.clients[name] = client
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
