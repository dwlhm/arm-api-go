package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Setup() (*gorm.DB, error) {

	dsn := os.Getenv("DB_LIVE")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db.AutoMigrate(&model.Calendar{}, &model.DetailJadwal{})

	return db, err

}
