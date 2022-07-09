package entity

import (
	"learn/common"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint      `gorm:"column:id;primaryKey,autoIncrement,notNull" json:"id,omitempty" mapstructure:"id"`
	Name     *string   `gorm:"column:name" json:"name,omitempty" mapstructer:"name"`
	Email    string    `gorm:"column:email;not null" json:"email,omitempty" mapstructer:"email"`
	Age      int       `gorm:"notNull" json:"age" mapstructer:"age"`
	Projects []Project `json:"projects,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" mapstructer:"projects"`
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
