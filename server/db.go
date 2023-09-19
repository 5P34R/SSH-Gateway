package server

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ContainerDB struct {
	gorm.Model
	ContainerID string
	UserCount   int
}

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("container.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&ContainerDB{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
