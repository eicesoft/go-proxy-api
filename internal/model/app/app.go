package app

//go:generate gormgen -structs App -input .
type App struct {
	Id        int32  `gorm:"primaryKey" json:"id"` // 主键
	Name      string `json:"name"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	Status    int8   `json:"status"`
	CreatedAt int64  `gorm:"time" json:"created_at"` // 创建时间
}

func (_ *App) TableName() string {
	return "apps"
}
