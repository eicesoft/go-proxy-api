///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package request_log

import (
	"fmt"

	"eicesoft/proxy-api/internal/model"
	"eicesoft/proxy-api/pkg/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *RequestLog {
	return new(RequestLog)
}

func NewQueryBuilder() *requestLogQueryBuilder {
	return new(requestLogQueryBuilder)
}

func (t *RequestLog) Assign(src interface{}) {
	core.StructCopy(t, src)
}

func (t *RequestLog) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

func (t *RequestLog) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (t *RequestLog) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Model(&RequestLog{}).Where("id = ?", t.Id).Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

type requestLogQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *requestLogQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = qb.buildUpdateQuery(db).Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *requestLogQueryBuilder) buildUpdateQuery(db *gorm.DB) *gorm.DB {
	ret := db.Model(&RequestLog{})
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	return ret
}

func (qb *requestLogQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *requestLogQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&RequestLog{}).Count(&c)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c = 0
	}
	return c, res.Error
}

func (qb *requestLogQueryBuilder) First(db *gorm.DB) (*RequestLog, error) {
	ret := &RequestLog{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *requestLogQueryBuilder) QueryOne(db *gorm.DB) (*RequestLog, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *requestLogQueryBuilder) QueryAll(db *gorm.DB) ([]*RequestLog, error) {
	var ret []*RequestLog
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *requestLogQueryBuilder) Limit(limit int) *requestLogQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *requestLogQueryBuilder) Offset(offset int) *requestLogQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *requestLogQueryBuilder) WhereId(p model.Predicate, value int32) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WhereIdIn(value []int32) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WhereIdNotIn(value []int32) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) OrderById(asc bool) *requestLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *requestLogQueryBuilder) WhereClientId(p model.Predicate, value int32) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "client_id", p),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WhereClientIdIn(value []int32) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "client_id", "IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WhereClientIdNotIn(value []int32) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "client_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) OrderByClientId(asc bool) *requestLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "client_id "+order)
	return qb
}

func (qb *requestLogQueryBuilder) WherePath(p model.Predicate, value string) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "path", p),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WherePathIn(value []string) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "path", "IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WherePathNotIn(value []string) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "path", "NOT IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) OrderByPath(asc bool) *requestLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "path "+order)
	return qb
}

func (qb *requestLogQueryBuilder) WhereParams(p model.Predicate, value string) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "params", p),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WhereParamsIn(value []string) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "params", "IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WhereParamsNotIn(value []string) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "params", "NOT IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) OrderByParams(asc bool) *requestLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "params "+order)
	return qb
}

func (qb *requestLogQueryBuilder) WhereAppId(p model.Predicate, value int32) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_id", p),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WhereAppIdIn(value []int32) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_id", "IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WhereAppIdNotIn(value []int32) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) OrderByAppId(asc bool) *requestLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "app_id "+order)
	return qb
}

func (qb *requestLogQueryBuilder) WhereCreatedAt(p model.Predicate, value int64) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WhereCreatedAtIn(value []int64) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) WhereCreatedAtNotIn(value []int64) *requestLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *requestLogQueryBuilder) OrderByCreatedAt(asc bool) *requestLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}
