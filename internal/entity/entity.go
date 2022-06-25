package entity

import (
	"errors"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          string
	StockId     *uint64
	Stock       Stock `gorm:"embedded"`
	Title       string
	Description string
	Price       float32
}

type Stock struct {
	gorm.Model
	ID        uint64
	ProductID string
	Count     uint64
}

func (invoice *Product) BeforeCreate(tx *gorm.DB) (err error) {
	invoice.ID, err = shortid.Generate()

	if err != nil {
		return errors.New("can't save invalid data")
	}
	return nil
}
