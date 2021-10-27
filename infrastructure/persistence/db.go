package persistence

import (
	"fmt"
	"github.com/chengziwj/mudan/domain/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories(user, pwd string, port uint16, host, name string) (*Repositories, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return &Repositories{User: NewUserRepository(db), db: db}, nil
}

func (s *Repositories) Close() error {
	return nil
}
