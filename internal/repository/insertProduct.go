package repository

import (
	"awsshop/internal/entity"

	"gorm.io/gorm"
)

func InsertProduct(db *gorm.DB, product *entity.Product) error {
	err := db.Create(&product).Error
	return err
}
