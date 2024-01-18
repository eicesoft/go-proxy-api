package router

import (
	"eicesoft/proxy-api/internal/controller/auth_controller"
	"eicesoft/proxy-api/internal/controller/check_controller"
	"eicesoft/proxy-api/internal/controller/order_controller"
	"eicesoft/proxy-api/pkg/mux"
)

// 设置Api路由
func setApiRouter(r *mux.Resource) {
	check_controller.New(r.GetLogger(), r.GetDb()).RegistryRouter(r)
	auth_controller.New(r.GetLogger(), r.GetDb()).RegistryRouter(r)
	order_controller.New(r.GetLogger(), r.GetDb()).RegistryRouter(r)
}
