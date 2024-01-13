package controller

import "eicesoft/proxy-api/pkg/core"

type RouteInfo struct {
	Path    string
	Closure core.HandlerFunc
}

type RouteInterface interface {
	Params() (string, core.HandlerFunc)
}

func (r *RouteInfo) Params() (string, core.HandlerFunc) {
	return r.Path, r.Closure
}
