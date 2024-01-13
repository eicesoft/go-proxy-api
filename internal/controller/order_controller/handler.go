package order_controller

import (
	"eicesoft/proxy-api/internal/controller"
	"eicesoft/proxy-api/internal/service/order_service"
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/db"
	"eicesoft/proxy-api/pkg/mux"
	"go.uber.org/zap"
)

const GroupRouterName = "/order"

var _ Handler = (*handler)(nil)

// Handler 用户控制器接口
type Handler interface {
	RegistryRouter(r *mux.Resource)
	List() *controller.RouteInfo
	Create() *controller.RouteInfo
	Update() *controller.RouteInfo
}

type handler struct {
	logger       *zap.Logger
	db           db.Repo
	orderService order_service.OrderService
}

func New(logger *zap.Logger, db db.Repo) Handler {
	return &handler{
		logger:       logger,
		db:           db,
		orderService: order_service.NewOrderService(db),
	}
}

func (h *handler) RegistryRouter(r *mux.Resource) {
	order := r.Mux.Group(GroupRouterName, core.WrapAuthHandler(r.Middles.Jwt))
	//typeOfA := reflect.TypeOf(h)
	//for i := 0; i <typeOfA.NumMethod(); i++ {
	//	f := typeOfA.Method(i)
	//	println(f.Name)
	//}
	order.GET(h.List().Params())
	order.POST(h.Create().Params())
	order.POST(h.Update().Params())
}
