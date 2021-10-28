package entity

import "github.com/chengziwj/moli/mcrypto"

type User struct {
	ID       uint32 `json:"id" gorm:"primary_key;auto_increment;column:uid"`
	Name     string `json:"name" gorm:"size:50;not null"`
	NickName string `json:"nick_name" gorm:"size:50;not null"`
	Pwd      string `json:"pwd" gorm:"size:255;not null"`
	Created  int    `json:"created" gorm:"not null"`
	Updated  int    `json:"updated" gorm:"not null"`
}

func (User) TableName() string  {
	return "t_user"
}

//BeforeSave gorm hook
func (u *User) BeforeSave() error {
	hashPwd, err := mcrypto.PasswordHash(u.Pwd)
	if err != nil {
		return err
	}
	u.Pwd = hashPwd
	return nil
}
