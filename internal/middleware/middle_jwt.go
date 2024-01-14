package middleware

import (
	"eicesoft/proxy-api/config"
	"eicesoft/proxy-api/internal/message"
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/errno"
	"eicesoft/proxy-api/pkg/token"
	"github.com/pkg/errors"
	"net/http"
)

func (m *middleware) Jwt(ctx core.Context) (userId int32, err errno.Error) {
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

	//m.logger.Info(fmt.Sprintf("expireAt: %d - %d = %d", claims.ExpiresAt, time.Now().Unix(),
	//	claims.ExpiresAt-time.Now().Unix()))
	//if claims.ExpiresAt-time.Now().Unix() <= 0 { //签名已过期
	//	err = errno.NewError(
	//		http.StatusUnauthorized,
	//		message.AuthorizationError,
	//		message.Text(message.AuthorizationError)).WithErr(errors.New("authorization is expired"))
	//
	//	return
	//}

	userId = claims.AppId
	if userId <= 0 {
		err = errno.NewError(
			http.StatusUnauthorized,
			message.AuthorizationError,
			message.Text(message.AuthorizationError)).WithErr(errors.New("claims.AppId <= 0 "))

		return
	}
	return
}
