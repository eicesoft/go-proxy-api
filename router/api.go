package router

import (
	"eicesoft/proxy-api/internal/controller/auth_controller"
	"eicesoft/proxy-api/internal/controller/order_controller"
	"eicesoft/proxy-api/internal/controller/user_controller"
	"eicesoft/proxy-api/pkg/mux"
)

// 设置Api路由
func setApiRouter(r *mux.Resource) {
	user_controller.New(r.GetLogger(), r.GetDb()).RegistryRouter(r)
	auth_controller.New(r.GetLogger(), r.GetDb()).RegistryRouter(r)
	order_controller.New(r.GetLogger(), r.GetDb()).RegistryRouter(r)
}
