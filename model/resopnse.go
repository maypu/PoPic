package model

type Response struct {
	Code    int
	Message string
	Result  string
}

// 构造函数
func NewResponse() *Response {
	response := Response{}
	response.Code = 200
	return &response
}
