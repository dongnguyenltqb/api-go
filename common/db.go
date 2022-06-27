package common

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var onceInitDB sync.Once

func GetDB() *gorm.DB {
	const url = "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	onceInitDB.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(url), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic(err)
		}
	})
	return db
}
