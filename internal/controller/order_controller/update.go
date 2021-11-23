package order_controller

import (
	"eicesoft/web-demo/internal/controller"
	"eicesoft/web-demo/internal/message"
	"eicesoft/web-demo/internal/service/order_service"
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/errno"
	"net/http"
)

type updateResponse struct {
	Id int32 `json:"id"` // 主键ID
}

func (h *handler) Update() *controller.RouteInfo {
	return &controller.RouteInfo{
		Path: "update",
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

			err := h.orderService.Update(c, req)
			if err != nil {
				panic(err)
			}

			res := new(updateResponse)
			res.Id = 1
			c.Payload(res)
		},
	}
}
