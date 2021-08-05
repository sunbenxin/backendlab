// Package db provides ...
package database

import (
	"gorm.io/gorm"
)

var db gorm.DB

func InitDb() {
	db, err := gorm.Open("", &gorm.Config{})
	if err != nil {
		return err
	}
}
