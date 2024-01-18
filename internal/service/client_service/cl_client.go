package client_service

import (
	"eicesoft/proxy-api/internal/model/consumption_log"
	"eicesoft/proxy-api/pkg/core"
	"eicesoft/proxy-api/pkg/db"
	"encoding/json"
	"log"
	"time"
)

var _ CLClientService = (*ClClientStruct)(nil)

type CheckResponse struct {
	ChargeCount  float64     `json:"chargeCount"`
	ChargeStatus float64     `json:"chargeStatus"`
	Code         string      `json:"code"`
	Data         interface{} `json:"data"`
	Message      string      `json:"message"`
}

type BlacklistResponse struct {
	ResultCode   string      `json:"resultCode"`
	ResultMsg    string      `json:"resultMsg"`
	RequestId    string      `json:"requestId"`
	ChargeCounts float64     `json:"chargeCounts"`
	ChargeStatus float64     `json:"chargeStatus"`
	Code         string      `json:"code"`
	ResultObj    interface{} `json:"resultObj"`
	Mobile       string      `json:"mobile"`
	Message      string      `json:"msg"`
	Forbid       string      `json:"forbid"`
	LuckyLevel   string      `json:"luckyLevel"`
}

type CLClientService interface {
	ClientService
	WithContext(context core.Context) *ClClientStruct
	RealCheck(params map[string]string) (CheckResponse, error)
	Check(params map[string]string) (CheckResponse, error)
	Blacklist(params map[string]string) (BlacklistResponse, error)
	addConsumptionLog(path string, chargeCount int32, params map[string]string, ratio int32)
}

type ClClientStruct struct {
	ClientStruct
}

func (c *ClClientStruct) addConsumptionLog(path string, chargeCount int32, params map[string]string, ratio int32) {
	buf, _ := json.Marshal(params)
	c.db.GetDbW().Create(&consumption_log.ConsumptionLog{
		BillingSum:   chargeCount * ratio,
		BillingCount: chargeCount,
		Path:         path,
		Params:       string(buf),
		AppId:        c.context.UserID(),
		ClientId:     ClientId,
		CreatedAt:    time.Now().Unix(),
	})
}

func (c *ClClientStruct) Check(params map[string]string) (CheckResponse, error) {
	resp, _ := c.Post("https://api.253.com/open/unn/batch-ucheck", params, "check")
	var response CheckResponse

	_ = ConvertStruct[CheckResponse](resp.(map[string]interface{}), &response)

	count := int32(response.ChargeCount)
	c.addConsumptionLog("check", count, params, 1)

	return response, nil
}

func (c *ClClientStruct) RealCheck(params map[string]string) (CheckResponse, error) {
	resp, _ := c.Post("https://api.253.com/open/mobstatus/mobstatus-query", params, "realcheck")
	log.Print(resp)
	var response CheckResponse

	_ = ConvertStruct[CheckResponse](resp.(map[string]interface{}), &response)

	c.addConsumptionLog("realcheck", 1, params, 2)

	return response, nil
}

func (c *ClClientStruct) Blacklist(params map[string]string) (BlacklistResponse, error) {
	resp, _ := c.Post("http://risk.253.com/risk_number/bforbid", params, "blacklist")
	log.Print(resp)
	var response BlacklistResponse

	_ = ConvertStruct[BlacklistResponse](resp.(map[string]interface{}), &response)

	count := int32(response.ChargeCounts)
	c.addConsumptionLog("blacklist", count, params, 1)

	return response, nil
}

func (c *ClClientStruct) WithContext(context core.Context) *ClClientStruct {
	c.context = context
	return c
}

func NewCLClientService(db db.Repo) *ClClientStruct {
	client := &ClClientStruct{
		ClientStruct: *NewClientService(db),
	}
	//client.restyClient.BaseURL = "https://api.253.com"

	return client
}
