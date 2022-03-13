package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mohamedhik/Golang-Book-Store/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connections()
	db = config.GetDB()
	db = AutoMigrate(&Book{})
}
