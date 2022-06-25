package repository

import (
	"awsshop/internal/entity"
	"awsshop/internal/productDto"

	"gorm.io/gorm"
)

func FindProductById(db *gorm.DB, id string) (productDto.ProductResponse, error) {
	var product entity.Product
	err := db.Preload("Stock").First(&product, "id = ?", id).Error
	dto := productDto.EntitytoResponsetDTO(&product)
	return dto, err
}
