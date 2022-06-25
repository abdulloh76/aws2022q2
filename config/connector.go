package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() (db *gorm.DB, err error) {
	dsn := "postgres://user:password@host:5432/db"

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
