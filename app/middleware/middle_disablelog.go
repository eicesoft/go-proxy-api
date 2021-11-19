package middleware

import "eicesoft/web-demo/pkg/core"

func (m *middleware) DisableLog() core.HandlerFunc {
	return func(c core.Context) {
		core.DisableTrace(c)
	}
}
