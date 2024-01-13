package router

import (
	"eicesoft/proxy-api/internal/middleware"
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/db"
	"eicesoft/proxy-api/pkg/mux"
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
