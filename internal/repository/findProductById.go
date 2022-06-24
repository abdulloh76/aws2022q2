package repository

import (
	"awsshop/internal/entity"

	"gorm.io/gorm"
)

func FindInvoiceById(db *gorm.DB, id string) (entity.Product, error) {
	var invoice entity.Product
	err := db.Preload("Stock").First(&invoice, "id = ?", id).Error
	return invoice, err
}
