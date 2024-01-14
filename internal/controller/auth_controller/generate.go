package auth_controller

import (
	"eicesoft/proxy-api/config"
	"eicesoft/proxy-api/internal/controller"
	"eicesoft/proxy-api/internal/message"
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/errno"
	"eicesoft/proxy-api/pkg/token"
	"net/http"
	"time"
)

type authRequest struct {
	AppKey    string `form:"app_key" binding:"required"`    // 用户名
	AppSecret string `form:"app_secret" binding:"required"` // 密码
}

type authResponse struct {
	Code int `json:"code"`
	Data struct {
		Authorization string `json:"authorization"` // 签名
		ExpireTime    int64  `json:"expire_time"`   // 过期时间
	} `json:"data"`
}

func (h *handler) Generate() *controller.RouteInfo {
	return &controller.RouteInfo{
		Path: "token",
		Closure: func(c core.Context) {
			req := new(authRequest)
			if err := c.ShouldBindQuery(req); err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					message.ParamBindError,
					message.Text(message.ParamBindError)).WithErr(err),
				)
				return
			}

			app := h.appService.Verification(req.AppKey, req.AppSecret)

			if app == nil || app.Status == 0 { //验证失败
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					message.GenerateTokenError,
					message.Text(message.GenerateTokenError)),
				)
				return
			}

			cfg := config.Get().JWT
			tokenString, err := token.New(cfg.Secret).JwtSign(app.Id, time.Hour*cfg.ExpireDuration)
			if err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					message.AuthorizationError,
					message.Text(message.AuthorizationError)).WithErr(err),
				)
				return
			}

			res := new(authResponse)
			res.Code = http.StatusOK
			res.Data.Authorization = tokenString
			res.Data.ExpireTime = time.Now().Add(time.Hour * cfg.ExpireDuration).Unix()

			c.Payload(res)

		},
	}
}
