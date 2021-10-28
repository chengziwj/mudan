package service

import (
	"github.com/chengziwj/mudan/domain/entity"
	"github.com/chengziwj/mudan/domain/repository"
)

var _ UserServiceInterface = &userService{}

type userService struct {
	us repository.UserRepository
}

type UserServiceInterface interface {
	SaveUser(user *entity.User) (*entity.User, error)
	GetUser(uid uint64)  (*entity.User, error)
}

func (u *userService) SaveUser(user *entity.User) (*entity.User, error) {
	return u.us.SaveUser(user)
}


func (u *userService) GetUser(uid uint64) (*entity.User, error) {
	return u.us.GetUser(uid)
}
