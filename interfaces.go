package notify

type IClient interface {
	I() IClient
	Name() string
	Send(message IMessage) IResult
}

// IMessage 消息
type IMessage interface {
	TransFormToRequestParams() any
}

// IResult 发送结果处理
type IResult interface {
	error
	// WithStatus 携带错误状态 RESULT_STATUS_ERROR 或 RESULT_STATUS_SUCCESS
	WithStatus(status int)
	// WithException  携带错误信息
	WithException(err error)
	// WithResult 携带请求接口返回值
	WithResult(response any)
	// Status 发送状态
	Status() int
	// Result 第三发返回数据
	Result() any
}
