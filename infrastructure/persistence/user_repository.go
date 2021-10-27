package persistence

import (
	"github.com/chengziwj/mudan/domain/entity"
	"github.com/chengziwj/mudan/domain/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

var _ repository.UserRepository = &UserRepository{}

func (r *UserRepository) SaveUser(user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
