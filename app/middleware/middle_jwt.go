package middleware

import (
	"eicesoft/web-demo/app/message"
	"eicesoft/web-demo/config"
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/errno"
	"eicesoft/web-demo/pkg/token"
	"github.com/pkg/errors"
	"net/http"
)

func (m *middleware) Jwt(ctx core.Context) (userId int64, err errno.Error) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		err = errno.NewError(
			http.StatusUnauthorized,
			message.AuthorizationError,
			message.Text(message.AuthorizationError)).WithErr(errors.New("Header 中缺少 Authorization 参数"))

		return
	}

	cfg := config.Get().JWT
	claims, errParse := token.New(cfg.Secret).JwtParse(auth)
	if errParse != nil {
		err = errno.NewError(
			http.StatusUnauthorized,
			message.AuthorizationError,
			message.Text(message.AuthorizationError)).WithErr(errParse)

		return
	}

	userId = claims.UserID
	if userId <= 0 {
		err = errno.NewError(
			http.StatusUnauthorized,
			message.AuthorizationError,
			message.Text(message.AuthorizationError)).WithErr(errors.New("claims.UserID <= 0 "))

		return
	}
	return
}
