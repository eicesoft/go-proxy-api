package order_service

import (
	"eicesoft/proxy-api/internal/model"
	orders "eicesoft/proxy-api/internal/model/order"
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/db"
)

var _ OrderService = (*orderService)(nil)

type OrderService interface {
	List(ctx core.Context) []*orders.Order
	Create(ctx core.Context, orderInfo *OrderInfo) (id int32, err error)
	Update(ctx core.Context, orderInfo *OrderInfo) (err error)
}

type orderService struct {
	db db.Repo
}

type OrderInfo struct {
	OrderNo      string `form:"order_no"`
	MemberId     int32  `form:"member_id"`
	CompanyId    int32  `form:"company_id"`
	CustomerName string `form:"customer_name"`
}

func NewOrderService(db db.Repo) *orderService {
	return &orderService{
		db: db,
	}
}

func (o *orderService) Update(ctx core.Context, orderInfo *OrderInfo) (err error) {
	orderModel := orders.NewModel()
	orderModel.OrderNo = orderInfo.OrderNo

	data := map[string]interface{}{
		"status": 100,
	}

	err = orders.NewQueryBuilder().
		WhereOrderNo(model.EqualPredicate, "123456").
		Updates(o.db.GetDbW().WithContext(ctx.RequestContext()), data)

	return
}

// Create 创建订单
func (o *orderService) Create(ctx core.Context, orderInfo *OrderInfo) (id int32, err error) {
	orderModel := orders.NewModel()
	orderModel.Assign(orderInfo)
	//core.StructCopy(orderModel, orderInfo)
	//orderModel.CompanyId = orderInfo.CompanyId
	//orderModel.OrderNo = orderInfo.OrderNo
	//orderModel.MemberId = orderInfo.MemberId
	//orderModel.CustomerName = orderInfo.CustomerName
	orderModel.DeviceInfo = "{}"

	id, err = orderModel.Create(o.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		panic(err)
	}
	return
}

func (o *orderService) List(ctx core.Context) []*orders.Order {
	orderList, err := orders.
		NewQueryBuilder().
		Limit(10).
		WhereIsClosed(model.EqualPredicate, 0).
		QueryAll(o.db.GetDbR().WithContext(ctx.RequestContext()))
	//
	//ctx.Logger().Info("call order service list",
	//	zap.Any("service", "orders"),
	//	zap.Any("data", orderList))

	if err != nil {
		ctx.Logger().Error(err.Error())
	}

	return orderList
}
