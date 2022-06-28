package repository

import (
	"awsshop/internal/entity"

	"gorm.io/gorm"
)

func FindProductById(db *gorm.DB, id string) (entity.Product, error) {
	var product entity.Product
	err := db.Preload("Product").Preload("Product.Stock").First(&product, "id = ?", id).Error
	return product, err
}
