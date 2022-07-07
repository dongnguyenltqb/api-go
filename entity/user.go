package entity

import (
	"learn/common"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint      `gorm:"column:id;primaryKey,autoIncrement,notNull" json:"id,omitempty"`
	Name     *string   `gorm:"column:name" json:"name,omitempty"`
	Email    string    `gorm:"column:email;not null" json:"email,omitempty"`
	Age      int       `gorm:"notNull" json:"age"`
	Projects []Project `json:"projects,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func CreateUser(u *User) *User {
	common.GetDB().Create(u)
	return u
}

func GetById(id int) *User {
	user := new(User)
	common.GetDB().Model(&User{}).First(user, map[string]interface {
	}{
		"id": id,
	})
	return user
}
