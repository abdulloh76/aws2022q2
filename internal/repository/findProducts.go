package repository

import (
	"awsshop/internal/entity"
	"awsshop/internal/productDto"

	"gorm.io/gorm"
)

func FindProducts(db *gorm.DB) ([]productDto.ProductsListResponse, error) {
	var products []productDto.ProductsListResponse
	err := db.Model(&entity.Product{}).Find(&products).Error
	return products, err
}
