package repository

import (
	"github.com/chengziwj/mudan/domain/entity"
)

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
}
