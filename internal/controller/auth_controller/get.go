package auth_controller

import (
	"eicesoft/proxy-api/config"
	"eicesoft/proxy-api/internal/message"
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/errno"
	"eicesoft/proxy-api/pkg/token"
	"math/rand"
	"net/http"
	"time"
)

type authRequest struct {
	UserName string `form:"username" binding:"required"` // 用户名
	Password string `form:"password" binding:"required"` // 密码
}

type authResponse struct {
	Code int `json:"code"`
	Data struct {
		Authorization string `json:"authorization"` // 签名
		ExpireTime    int64  `json:"expire_time"`   // 过期时间
	} `json:"data"`
}

func (h *handler) Get() (string, core.HandlerFunc) {
	return "get", func(c core.Context) {
		req := new(authRequest)
		//res := new(detailResponse)
		if err := c.ShouldBindQuery(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.ParamBindError,
				message.Text(message.ParamBindError)).WithErr(err),
			)
			return
		}

		cfg := config.Get().JWT
		tokenString, err := token.New(cfg.Secret).JwtSign(rand.Int31(), time.Hour*cfg.ExpireDuration)
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
	}
}
