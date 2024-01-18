package check_controller

import (
	"eicesoft/proxy-api/internal/service/client_service"
	"eicesoft/proxy-api/internal/service/user_service"
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/db"
	"eicesoft/proxy-api/pkg/mux"
	"go.uber.org/zap"
)

const GroupRouterName = "/check"

var _ Handler = (*handler)(nil)

// Handler 用户控制器接口
type Handler interface {
	RegistryRouter(r *mux.Resource)
	CheckNumbers() (string, core.HandlerFunc)
	Blacklist() (string, core.HandlerFunc)
}

type handler struct {
	logger        *zap.Logger
	db            db.Repo
	userService   user_service.UserService
	clientService client_service.CLClientService
}

func New(logger *zap.Logger, db db.Repo) Handler {
	return &handler{
		logger:        logger,
		db:            db,
		userService:   user_service.NewUserService(db),
		clientService: client_service.NewCLClientService(db),
	}
}

func (h *handler) RegistryRouter(r *mux.Resource) {
	user := r.Mux.Group(GroupRouterName, core.WrapAuthHandler(r.Middles.Jwt))

	user.GET(h.CheckNumbers())
	user.GET(h.RealCheckNumber())
	user.GET(h.Blacklist())
}
