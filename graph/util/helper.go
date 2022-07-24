package util

import "gorm.io/gorm"

func FuncOrder(db *gorm.DB) *gorm.DB {
	return db.Order("\"order\"")
}