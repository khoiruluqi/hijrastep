package db

import (
	"log"

	"hijrastep/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbURL string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&models.Bab{},
		&models.SubBab{},
		&models.Content{},
		&models.Material{},
		&models.Test{},
		&models.Question{},
		&models.ContentLog{},
		&models.ExamLog{},
	)

	return db
}
