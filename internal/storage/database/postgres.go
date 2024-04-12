package database

import (
	"log"
	"os"

	"github.com/loveavoider/avito-banners/internal/repository/banner/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgres() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Banner{}, &entity.Tag{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
