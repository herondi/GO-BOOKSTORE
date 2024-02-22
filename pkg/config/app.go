package config

import(
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)


var (
	DB *gorm.DB
)

func Connect(){
d, err := gorm.Open("mysql", &mysql.Config{
Host: "localhost",
User: "root",
}