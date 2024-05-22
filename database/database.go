package database

import (
	"os"

	"github.com/hosseinmirzapur/ecombot/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func RegisterDB() error {
	conn, err := gorm.Open(
		mysql.Open(os.Getenv("DB_URI")),
		&gorm.Config{},
	)
	if err != nil {
		return err
	}

	sqlDB, err := conn.DB()
	if err != nil {
		return err
	}

	if err = sqlDB.Ping(); err != nil {
		return err
	}

	db = conn
	return nil
}

func AutoMigrate() error {
	return db.AutoMigrate(
		&models.Color{},
		&models.Image{},
		&models.Product{},
		&models.User{},
		&models.Video{},
	)
}

func DB() *gorm.DB {
	return db
}
