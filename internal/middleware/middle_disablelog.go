package middleware

import "eicesoft/proxy-api/pkg/core"

func (m *middleware) DisableLog() core.HandlerFunc {
	return func(c core.Context) {
		core.DisableTrace(c)
	}
}
