package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *EnvConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
