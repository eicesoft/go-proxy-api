package db

import (
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/time_parse"
	"eicesoft/proxy-api/pkg/trace"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"time"
)

const (
	callBackBeforeName = "core:before"
	callBackAfterName  = "core:after"
	startTime          = "_start_time"
)

type TracePlugin struct{}

func (op *TracePlugin) Name() string {
	return "tracePlugin"
}

func (op *TracePlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前
	_ = db.Callback().Create().Before("gorm:create").Register(callBackBeforeName, before)
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	_ = db.Callback().Delete().Before("gorm:delete").Register(callBackBeforeName, before)
	_ = db.Callback().Update().Before("gorm:update").Register(callBackBeforeName, before)

	// 结束后
	_ = db.Callback().Create().After("gorm:create").Register(callBackAfterName, after)
	_ = db.Callback().Query().After("gorm:query").Register(callBackAfterName, after)
	_ = db.Callback().Delete().After("gorm:delete").Register(callBackAfterName, after)
	_ = db.Callback().Update().After("gorm:update").Register(callBackAfterName, after)
	return
}

var _ gorm.Plugin = &TracePlugin{}

func before(db *gorm.DB) {
	db.InstanceSet(startTime, time.Now())
	return
}

func after(db *gorm.DB) {
	_ctx := db.Statement.Context
	ctx, ok := _ctx.(core.StdContext)
	if !ok {
		return
	}

	_ts, isExist := db.InstanceGet(startTime)
	if !isExist {
		return
	}

	ts, ok := _ts.(time.Time)
	if !ok {
		return
	}

	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)

	sqlInfo := new(trace.SQL)
	sqlInfo.Timestamp = time_parse.CSTLayoutString()
	sqlInfo.SQL = sql
	sqlInfo.Stack = utils.FileWithLineNum()
	sqlInfo.Rows = db.Statement.RowsAffected
	sqlInfo.CostSeconds = time.Since(ts).Seconds()
	if ctx.Trace != nil {
		ctx.Trace.AppendSQL(sqlInfo)
	}

	return
}
