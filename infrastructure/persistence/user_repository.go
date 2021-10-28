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

func (r *UserRepository) GetUser(uid uint64) (*entity.User, error) {
	var user = &entity.User{}
	err := r.db.First(user, uid).Error
	return user, err
}
func (r *UserRepository) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}
