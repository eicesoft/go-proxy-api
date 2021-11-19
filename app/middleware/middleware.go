package middleware

import (
	"eicesoft/web-demo/pkg/core"
	"go.uber.org/zap"
)

var _ Middleware = (*middleware)(nil)

type Middleware interface {
	// i 为了避免被其他包实现
	i()

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

func (m *middleware) i() {}
