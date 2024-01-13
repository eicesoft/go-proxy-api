package request_log

//go:generate gormgen -structs RequestLog -input .
type RequestLog struct {
	Id        int32  `gorm:"primaryKey" json:"id"` // 主键
	ClientId  int32  `json:"client_id"`
	Path      string `json:"path"`
	Params    string `json:"params"`
	AppId     int32  `json:"app_id"`
	CreatedAt int64  `gorm:"time" json:"created_at"` // 创建时间
}

func (_ *RequestLog) TableName() string {
	return "request_logs"
}
