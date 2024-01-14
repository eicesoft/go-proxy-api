package errno

import (
	"eicesoft/proxy-api/internal/message"
	"fmt"
	"github.com/pkg/errors"
)

var _ Error = (*err)(nil)

type Error interface {
	p()
	// WithErr 设置错误信息
	WithErr(err error) Error
	// GetBusinessCode 获取 Business Code
	GetBusinessCode() int
	// GetHttpCode 获取 HTTP Code
	GetHttpCode() int
	// GetMsg 获取 Msg
	GetMsg() string
	// ToString 返回 JSON 格式的错误详情
	GetJson() *message.Failure
}

type err struct {
	HttpCode     int
	BusinessCode int
	Message      string
	Err          error
}

func NewError(httpCode, businessCode int, msg string) Error {
	return &err{
		HttpCode:     httpCode,
		BusinessCode: businessCode,
		Message:      msg,
	}
}

func (e *err) p() {}

func (e *err) WithErr(err error) Error {
	e.Err = errors.WithStack(err)
	return e
}

func (e *err) GetHttpCode() int {
	return e.HttpCode
}

func (e *err) GetBusinessCode() int {
	return e.BusinessCode
}

func (e *err) GetMsg() string {
	return e.Message
}

// GetJson 返回错误详情
func (e *err) GetJson() *message.Failure {
	var err message.Failure
	if e.Err != nil {
		err = message.Failure{
			Code:    e.BusinessCode,
			Data:    e.Err,
			Message: fmt.Sprintf("%s: %s", e.Message, e.Err.Error()),
		}
	} else {
		err = message.Failure{
			Code:    e.BusinessCode,
			Data:    e.Err,
			Message: fmt.Sprintf("%s", e.Message),
		}
	}

	return &err
}
