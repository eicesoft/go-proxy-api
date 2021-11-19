package user_controller

import (
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/db"
	"eicesoft/web-demo/pkg/mux"
	"go.uber.org/zap"
)

const GroupRouterName = "/user"

var _ Handler = (*handler)(nil)

// Handler 用户控制器接口
type Handler interface {
	RegistryRouter(r *mux.Resource)
	Detail() (string, core.HandlerFunc)
	Test() (string, core.HandlerFunc)
}

type handler struct {
	logger *zap.Logger
	db     db.Repo
}

func New(logger *zap.Logger, db db.Repo) Handler {
	return &handler{
		logger: logger,
		db:     db,
	}
}

func (h *handler) RegistryRouter(r *mux.Resource) {
	user := r.Mux.Group(GroupRouterName)

	user.GET(h.Detail())
	user.GET(h.Test())
}
