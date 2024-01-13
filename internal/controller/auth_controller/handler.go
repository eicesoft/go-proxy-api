package auth_controller

import (
	"eicesoft/proxy-api/internal/controller"
	"eicesoft/proxy-api/internal/service"
	"eicesoft/proxy-api/pkg/db"
	"eicesoft/proxy-api/pkg/mux"
	"go.uber.org/zap"
)

const GroupRouterName = "/auth"

var _ Handler = (*handler)(nil)

// Handler 用户控制器接口
type Handler interface {
	RegistryRouter(r *mux.Resource)
	Generate() *controller.RouteInfo
}

type handler struct {
	logger     *zap.Logger
	db         db.Repo
	appService service.AppService
}

func New(logger *zap.Logger, db db.Repo) Handler {
	return &handler{
		logger:     logger,
		appService: service.NewAppService(db),
		db:         db,
	}
}

func (h *handler) RegistryRouter(r *mux.Resource) {
	auth := r.Mux.Group(GroupRouterName)
	auth.GET(h.Generate().Params())
}
