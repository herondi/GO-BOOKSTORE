package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

// Connect initializes the database connection
func Connect() {
	d, err := gorm.Open("mysql", "root@localhost/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = d
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}