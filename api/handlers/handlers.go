package handlers

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func InitHandlers(database *gorm.DB) {
	db = database
}
