package message

import "fmt"

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	/* 服务级错误码 */
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	GenerateTokenError = 10105
	CallHTTPError      = 10110
)

var codeText = map[int]string{
	ServerError:        "Internal Server Error",
	TooManyRequests:    "Too Many Requests",
	ParamBindError:     "参数信息有误",
	AuthorizationError: "签名信息有误",
	GenerateTokenError: "生成Token信息有误",
	CallHTTPError:      "调用第三方 HTTP 接口失败 %s",
}

func Text(code int, a ...any) string {
	return fmt.Sprintf(codeText[code], a...)
}
