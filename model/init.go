package model

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database



func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(20000)
	db.DB().SetMaxIdleConns(0)
}