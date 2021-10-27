package entity

import "github.com/chengziwj/moli/mcrypto"

type User struct {
	ID       uint32 `json:"id" gorm:"primary_key;auto_increment"`
	Name     string `json:"name" gorm:"size:50;not null"`
	NickName string `json:"nick_name" gorm:"size:50;not null"`
	Pwd      string `json:"pwd" gorm:"size:255;not null"`
	Created  int    `json:"created" gorm:"created"`
	Updated  int    `json:"updated" gorm:"updated"`
}

func (u *User) BeforeSave() error {
	hashPwd, err := mcrypto.PasswordHash(u.Pwd)
	if err != nil {
		return err
	}
	u.Pwd = hashPwd
	return nil
}
