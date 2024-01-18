package check_controller

import (
	"eicesoft/proxy-api/internal/message"
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/errno"
	"net/http"
)

type checkReq struct {
	Mobiles string `form:"mobiles"`
}

type blacklistReq struct {
	Mobiles     string `form:"mobiles"`
	ForbidLevel string `form:"forbidLevel"`
}

type realCheckReq struct {
	Mobile string `form:"mobile"`
}

//type checkResult struct {
//	Area          string `json:"area"`
//	ChargesStatus string `json:"chargesStatus"`
//	Mobile        string `json:"mobile"`
//	NumberType    string `json:"numberType"`
//	Status        string `json:"status"`
//}

func (h *handler) CheckNumbers() (string, core.HandlerFunc) {
	return "numbers", func(c core.Context) {
		req := new(checkReq)

		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.ParamBindError,
				message.Text(message.ParamBindError)).WithErr(err),
			)
			return
		}

		resp, err := h.clientService.
			WithContext(c).
			Check(map[string]string{"mobiles": req.Mobiles})

		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.CallHTTPError,
				message.Text(message.CallHTTPError, err.Error())).WithErr(err),
			)
		}

		c.Success(200, "", resp.Data)
	}
}

func (h *handler) RealCheckNumber() (string, core.HandlerFunc) {
	return "realtime", func(c core.Context) {
		req := new(realCheckReq)

		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.ParamBindError,
				message.Text(message.ParamBindError)).WithErr(err),
			)
			return
		}

		resp, err := h.clientService.
			WithContext(c).
			RealCheck(map[string]string{"mobile": req.Mobile})

		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.CallHTTPError,
				message.Text(message.CallHTTPError, err.Error())).WithErr(err),
			)
		}

		c.Success(200, "", resp.Data)
	}
}

func (h *handler) Blacklist() (string, core.HandlerFunc) {
	return "blacklist", func(c core.Context) {
		req := new(blacklistReq)

		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.ParamBindError,
				message.Text(message.ParamBindError)).WithErr(err),
			)
			return
		}

		resp, err := h.clientService.
			WithContext(c).
			Blacklist(map[string]string{"mobiles": req.Mobiles, "forbidLevel": req.ForbidLevel})

		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.CallHTTPError,
				message.Text(message.CallHTTPError, err.Error())).WithErr(err),
			)
		}

		c.Success(200, "", resp.ResultObj)
	}
}
