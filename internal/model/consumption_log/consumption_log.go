package consumption_log

//go:generate gormgen -structs ConsumptionLog -input .
type ConsumptionLog struct {
	Id           int64  `json:"id"`            //
	BillingCount int32  `json:"billing_count"` //
	BillingSum   int32  `json:"billing_sum"`   //
	Path         string `json:"path"`          //
	Params       string `json:"params"`
	AppId        int32  `json:"app_id"`     //
	ClientId     int32  `json:"client_id"`  //
	CreatedAt    int64  `json:"created_at"` //
}

func (_ *ConsumptionLog) TableName() string {
	return "consumption_logs"
}
