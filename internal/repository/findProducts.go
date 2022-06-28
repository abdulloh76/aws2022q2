package repository

import (
	"awsshop/internal/entity"
	"awsshop/internal/productDto"

	"gorm.io/gorm"
)

func FindProducts(db *gorm.DB) ([]productDto.ProductsListResponse, error) {
	var products []productDto.ProductsListResponse
	// err := db.Joins("Stock").Find(&products).Error
	err := db.Model(&entity.Product{}).Select("products.id, products.title, products.description, products.price, stocks.count").Joins("inner join stocks on products.id=stocks.product_id").Find(&products).Error
	return products, err
}
