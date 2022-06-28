package entity

import (
	"errors"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          string
	Stock       Stock `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title       string
	Description string
	Price       float32
}

type Stock struct {
	gorm.Model
	ID        uint64
	Count     uint64
	ProductID string
}

func (invoice *Product) BeforeCreate(tx *gorm.DB) (err error) {
	invoice.ID, err = shortid.Generate()

	if err != nil {
		return errors.New("can't save invalid data")
	}
	return nil
}
