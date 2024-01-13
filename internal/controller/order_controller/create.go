package order_controller

import (
	"eicesoft/proxy-api/internal/controller"
	"eicesoft/proxy-api/internal/message"
	"eicesoft/proxy-api/internal/service/order_service"
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/errno"
	"net/http"
)

type createResponse struct {
	Id int32 `json:"id"` // 主键ID
}

func (h *handler) Create() *controller.RouteInfo {
	return &controller.RouteInfo{
		Path: "create",
		Closure: func(c core.Context) {
			req := new(order_service.OrderInfo)
			if err := c.ShouldBindForm(req); err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					message.ParamBindError,
					message.Text(message.ParamBindError)).WithErr(err),
				)
				return
			}

			id, _ := h.orderService.Create(c, req)

			res := new(createResponse)
			res.Id = id
			c.Payload(res)
		},
	}
}
