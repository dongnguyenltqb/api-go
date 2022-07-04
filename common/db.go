package common

import (
	"errors"
	"learn/config"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var onceInitDB sync.Once

func GetDB() *gorm.DB {
	dsn := config.Get().DSN
	if dsn == "" {
		panic(errors.New("db_dsn must be valid"))
	}
	onceInitDB.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic(err)
		}
	})
	return db
}
