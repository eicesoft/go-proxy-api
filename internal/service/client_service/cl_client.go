package client_service

import (
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/db"
)

var _ CLClientService = (*clClientStruct)(nil)

type CLClientService interface {
	ClientService
	WithContext(context core.Context) *clClientStruct

	Test(params map[string]string) (interface{}, error)
	Test2(params map[string]string) (interface{}, error)
}

type clClientStruct struct {
	ClientStruct
}

func (c *clClientStruct) WithContext(context core.Context) *clClientStruct {
	c.context = context
	return c
}

func (c *clClientStruct) Test(params map[string]string) (interface{}, error) {
	resp, _ := c.Post("/post", params)
	return resp, nil
}

func (c *clClientStruct) Test2(params map[string]string) (interface{}, error) {
	resp, _ := c.Get("/get", params)
	return resp, nil
}

func NewCLClientService(db db.Repo) *clClientStruct {
	client := &clClientStruct{
		ClientStruct: *NewClientService(db),
	}
	client.restyClient.BaseURL = "https://httpbin.org"

	return client
}
