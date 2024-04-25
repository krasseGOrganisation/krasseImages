package db

import (
	"os"
	"sample-api/db/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db

	DB.Exec("CREATE DATABASE IF NOT EXISTS sample")
	DB.Exec("USE sample")

	DB.AutoMigrate(
		&entities.Entry{},
	)
}
