package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/tree/master/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connections() {
	d, err := gorm.Open("mysql", "dhik:Dhik@2000/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

query := ("
	SELECT username FROM table WHERE passwor
");

