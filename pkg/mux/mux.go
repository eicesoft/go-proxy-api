package mux

import (
	"eicesoft/web-demo/internal/middleware"
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/db"
	"go.uber.org/zap"
)

var _ ResourceInterface = (*Resource)(nil)

type Resource struct {
	Mux     core.Mux
	Middles middleware.Middleware
	logger  *zap.Logger
	db      db.Repo
}

type ResourceInterface interface {
	SetDb(db db.Repo)
	GetDb() db.Repo

	SetLogger(logger *zap.Logger)
	GetLogger() *zap.Logger
}

func (r *Resource) SetDb(db db.Repo) {
	r.db = db
}

func (r *Resource) GetDb() db.Repo {
	return r.db
}

func (r *Resource) SetLogger(logger *zap.Logger) {
	r.logger = logger
}

func (r *Resource) GetLogger() *zap.Logger {
	return r.logger
}
