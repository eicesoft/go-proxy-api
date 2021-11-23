package router

import (
	"eicesoft/web-demo/internal/middleware"
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/db"
	"eicesoft/web-demo/pkg/mux"
	"go.uber.org/zap"
)

func InitMux(logger *zap.Logger, db db.Repo) (core.Mux, error) {
	m, err := core.New(logger)
	if err != nil {
		panic(err)
	}

	r := new(mux.Resource)
	r.Mux = m
	r.SetDb(db)
	r.SetLogger(logger)
	r.Middles = middleware.New(logger)
	//r.RegistryRouters()
	setApiRouter(r)

	return m, nil
}
