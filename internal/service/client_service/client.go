package client_service

import (
	"eicesoft/proxy-api/internal/model/request_log"
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/db"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"reflect"
	"time"
)

const (
	AgentName = "Proxy API 0.9.1"
	ClientId  = 1
)

var _ ClientService = (*ClientStruct)(nil)

func ConvertStruct[T any](input map[string]interface{}, output *T) error {
	outputVal := reflect.ValueOf(output)
	outputValue := outputVal.Elem()
	outputType := outputValue.Type()

	for i := 0; i < outputValue.NumField(); i++ {
		fieldName := outputType.Field(i).Tag.Get("json")
		mapValue, ok := input[fieldName]

		if !ok {
			continue // Field not found in the map
		}

		fieldValue := outputValue.Field(i)
		//fieldType := fieldValue.Type()
		//if fieldType != reflect.TypeOf(mapValue) {
		//	convertedValue := reflect.ValueOf(fmt.Sprintf("%v", mapValue)).Convert(fieldType)
		//	fieldValue.Set(convertedValue)
		//} else {
		if mapValue != nil {
			fieldValue.Set(reflect.ValueOf(mapValue))
		}
		//}
	}

	return nil
}

type ClientService interface {
	h() // private 为了避免被其他包实现
	Post(url string, params map[string]string, path string) (interface{}, error)
	Get(url string, params map[string]string, path string) (interface{}, error)
}

type ClientStruct struct {
	restyClient *resty.Client
	db          db.Repo
	context     core.Context
}

// Get request data to url
func (c *ClientStruct) Get(
	url string,
	params map[string]string,
	path string,
) (interface{}, error) {
	t := new(interface{})

	resp, err := c.restyClient.R().
		SetResult(t).
		SetQueryParams(params).
		SetHeader("User-Agent", AgentName).
		ForceContentType("application/json").
		Get(url)

	if err != nil && resp.IsSuccess() {
		return *t, err
	}

	return *t, nil
}

// Post request data to url
func (c *ClientStruct) Post(url string, params map[string]string, path string) (interface{}, error) {
	buf, _ := json.Marshal(params)

	c.db.GetDbW().Create(&request_log.RequestLog{
		ClientId:  ClientId,
		Path:      path,
		Params:    string(buf),
		AppId:     c.context.UserID(),
		CreatedAt: time.Now().Unix(),
	})

	params["appId"] = "Kssn4per"
	params["appKey"] = "KIzwjxaU"

	t := new(interface{})
	resp, err := c.restyClient.R().
		SetFormData(params).
		SetResult(t).
		SetHeader("User-Agent", AgentName).
		ForceContentType("application/json").
		Post(url)

	if err != nil && resp.IsSuccess() {
		return *t, err
	}

	return *t, nil
}

func (c *ClientStruct) h() {}

func NewClientService(db db.Repo) *ClientStruct {
	client := &ClientStruct{
		restyClient: resty.New(),
		db:          db,
	}

	return client
}
