package core

import (
	modelV1 "fuge/app/models/v1"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(c *ConfYaml) *gorm.DB {
	var err error
	DB, err = gorm.Open(postgres.Open(c.DBDSN), &gorm.Config{})
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(30)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(modelV1.Models...)
	return DB
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}
