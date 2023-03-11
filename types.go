package notify

type IClient interface {
	I() IClient
	Send(message IMessage) (result Result, err error)
}

type IMessage interface {
	// I 初始化并返回IMessage
	I() IMessage
	// AtMobiles 消息需要@人员手机号
	AtMobiles(mobiles []string)
	// AtUserIds 消息需要@人员userid
	AtUserIds(userIds []string)
	// AtAll @所有人
	AtAll()
	// ToJson 转json字符串
	ToJson() string
}
