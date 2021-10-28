package repository

import (
	"github.com/chengziwj/mudan/domain/entity"
)

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
	GetUser(uint642 uint64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
}
