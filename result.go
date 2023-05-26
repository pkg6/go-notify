package notify

type Result struct {
	client    IClient
	message   IMessage
	status    int
	exception error
	result    any
}

// BuildResult 在网关快速生成结构体
func BuildResult(client IClient, message IMessage) IResult {
	rt := &Result{client: client, message: message}
	rt.WithStatus(RESULT_STATUS_SUCCESS)
	return rt
}

// WithResult 携带第三方相应数据
func (r *Result) WithResult(response any) {
	r.result = response
}

// WithStatus 携带状态值
func (r *Result) WithStatus(status int) {
	r.status = status
}

// WithException 异常
func (r *Result) WithException(err error) {
	if err != nil {
		r.WithStatus(RESULT_STATUS_ERROR)
	}
	r.exception = err
}

// Status 消息发送状态
func (r *Result) Status() int {
	return r.status
}

// Result WithException 异常
func (r *Result) Result() any {
	return r.result
}

//返回错误信息
func (r *Result) Error() string {
	if r.exception == nil {
		return ""
	}
	return r.exception.Error()
}
