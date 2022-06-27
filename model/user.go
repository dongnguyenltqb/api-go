package model

import (
	"learn/common"
)

type User struct {
	ID       uint    `gorm:"column:id;primaryKey,autoIncrement,notNull"`
	Name     *string `gorm:"column:name"`
	Email    string  `gorm:"column:email;not null"`
	Age      uint8   `gorm:"notNull"`
	Projects []Project
}

func Create(u *User) *User {
	common.GetDB().Create(u)
	return u
}
