package user_controller

import (
	"eicesoft/web-demo/internal/message"
	"eicesoft/web-demo/internal/service/client_service"
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type detailRequest struct {
	UserName string `uri:"username"` // 用户名
}

type detailResponse struct {
	Id       int32       `json:"id"`        // 用户主键ID
	UserName string      `json:"user_name"` // 用户名
	NickName string      `json:"nick_name"` // 昵称
	Data     interface{} `json:"data"`
}

type testResponse struct {
	Args    interface{} `json:"args"`
	Data    string      `json:"data"`
	Files   interface{} `json:"files"`
	Form    interface{} `json:"form"`
	Headers interface{} `json:"headers"`
	Json    interface{} `json:"json"`
	Origin  string      `json:"origin"`
	Url     string      `json:"url"`
}

func (h *handler) Test() (string, core.HandlerFunc) {
	return "test", func(c core.Context) {
		//h.clientService.
		resp, err := h.clientService.
			WithContext(c).
			Test(map[string]string{"username": "testuser", "password": "dsgsdg"})

		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.CallHTTPError,
				message.Text(message.CallHTTPError, err.Error())).WithErr(err),
			)
		}

		var response testResponse

		_ = client_service.ConvertStruct[testResponse](resp.(map[string]interface{}), &response)

		c.Payload(gin.H{
			"message": "这个是一个Gin.H消息",
			"data":    response,
		})
	}
}

// Detail 用户详情
// @Summary 用户详情
// @Description 用户详情
// @Tags User
// @Accept  json
// @Produce  json
// @Param username path string true "用户名"
// @Success 200 {object} detailResponse
// @Failure 400 {object} message.Failure
// @Failure 401 {object} message.Failure
// @Router /user/get/{username} [get]
func (h *handler) Detail() (string, core.HandlerFunc) {
	return "get/:username", func(c core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.ParamBindError,
				message.Text(message.ParamBindError)).WithErr(err),
			)
			return
		}

		u := h.userService.Get()

		//u := user.User{}
		//h.db.GetDbR().WithContext(c.RequestContext()).First(&u)

		if req.UserName != "sdg" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.ParamBindError,
				message.Text(message.ParamBindError)),
			)
			return
		}

		res.Id = c.UserID()
		res.UserName = req.UserName
		res.NickName = req.UserName + "_nick"
		res.Data = u

		c.Payload(res)
	}
}
