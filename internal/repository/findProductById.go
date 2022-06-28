package repository

import (
	"awsshop/internal/entity"
	"awsshop/internal/productDto"

	"gorm.io/gorm"
)

func FindProductById(db *gorm.DB, id string) (productDto.ProductResponse, error) {
	var product productDto.ProductResponse
	err := db.Model(&entity.Product{}).Select("products.id, products.title, products.description, products.price, stocks.count").Joins("inner join stocks on products.id=stocks.product_id").First(&product, "products.id = ?", id).Error
	return product, err
}
