package router

import (
	"eicesoft/web-demo/app/controller/auth_controller"
	"eicesoft/web-demo/app/controller/user_controller"
	"eicesoft/web-demo/pkg/mux"
)

// 设置Api路由
func setApiRouter(r *mux.Resource) {
	user_controller.New(r.GetLogger(), r.GetDb()).RegistryRouter(r)
	auth_controller.New(r.GetLogger(), r.GetDb()).RegistryRouter(r)
}
