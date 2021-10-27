package service

import (
	"github.com/chengziwj/mudan/domain/entity"
	"github.com/chengziwj/mudan/domain/repository"
)

type userService struct {
	us repository.UserRepository
}

type UserServiceInterface interface {
	SaveUser(user *entity.User) (*entity.User, error)
}

func (u *userService) SaveUser(user *entity.User) (*entity.User, error) {
	return u.us.SaveUser(user)
}
