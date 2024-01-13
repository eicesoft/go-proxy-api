package service

import (
	"eicesoft/proxy-api/internal/model"
	"eicesoft/proxy-api/internal/model/app"
	"eicesoft/proxy-api/pkg/db"
)

var _ AppService = (*appService)(nil)

type AppService interface {
	p() // private 为了避免被其他包实现
	Verification(appKey string, appSecret string) *app.App
}

type appService struct {
	db db.Repo
}

func (h *appService) Verification(appKey string, appSecret string) *app.App {
	a := &app.App{}
	a, _ = app.NewQueryBuilder().
		WhereAppKey(model.EqualPredicate, appKey).
		WhereAppSecret(model.EqualPredicate, appSecret).
		QueryOne(h.db.GetDbR())

	return a
}

func (h *appService) p() {}

func NewAppService(db db.Repo) *appService {
	return &appService{
		db: db,
	}
}
