package entity

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type ProjectAttribute struct {
	Tag *string `json:"tag"`
}

type Project struct {
	gorm.Model
	ID        uint              `gorm:"column:id;primarykey,autoIncrement"`
	Name      *string           `gorm:"column:name"`
	Attribute *ProjectAttribute `gorm:"column:attribute"`
	UserId    *uint             `gorm:"column:user_id"`
	User      *User             `gorm:"foreignKey:user_id"`
}

// value là data từ database, func này lấy giá trị từ value và gắn vào attr
func (attr *ProjectAttribute) Scan(value interface{}) error {
	if attr != nil {
		var result ProjectAttribute
		err := json.Unmarshal(value.([]byte), &result)
		*attr = result
		return err
	}
	return nil
}

// generate giá trị từ struct để đẩy vào db
// với struct thì định dạng là []byte
func (attr *ProjectAttribute) Value() (driver.Value, error) {
	if attr == nil {
		return nil, nil
	}
	return json.Marshal(attr)
}
