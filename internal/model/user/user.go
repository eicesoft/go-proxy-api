package user

//go:generate gormgen -structs User -input .
type User struct {
	Id        int32  `gorm:"primaryKey" json:"id"` // 主键
	Uid       string `json:"uid"`
	Type      int8   `json:"type"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	IsDelete  int8   `json:"is_delete"`
	CreatedAt int32  `gorm:"time" json:"created_at"` // 创建时间
	UpdatedAt int32  `gorm:"time" json:"updated_at"` // 更新时间
}

func (_ *User) TableName() string {
	return "members"
}
