package user_service

import (
	"eicesoft/web-demo/internal/model/user"
	"eicesoft/web-demo/pkg/db"
)

var _ UserService = (*userService)(nil)

type UserService interface {
	// private 为了避免被其他包实现
	p()
	Get() user.User
}

type userService struct {
	db db.Repo
}

func (us *userService) Get() user.User {
	user := user.User{}
	user.Uid = "200"
	user.Email = "Email@aefasf.com"
	//user.Create(us.db.GetDbR().WithContext(ctx))
	us.db.GetDbR().First(&user)

	return user
}

func NewUserService(db db.Repo) *userService {
	return &userService{
		db: db,
	}
}

func (us *userService) p() {}
