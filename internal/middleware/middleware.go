package middleware

import (
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/errno"
	"go.uber.org/zap"
)

var _ Middleware = (*middleware)(nil)

type Middleware interface {
	// private 为了避免被其他包实现
	p()

	// Jwt JWT 中间件
	Jwt(ctx core.Context) (userId int32, err errno.Error)

	// DisableLog 不记录日志
	DisableLog() core.HandlerFunc
}

type middleware struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *middleware {
	return &middleware{
		logger: logger,
	}
}

func (m *middleware) p() {}
