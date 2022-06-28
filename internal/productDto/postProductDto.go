package productDto

import "awsshop/internal/entity"

type ProductBody struct {
	Count       uint64
	Title       string
	Description string
	Price       float32
}

func RequestDTOtoEntity(dto *ProductBody) entity.Product {
	product := entity.Product{
		Title:       dto.Title,
		Description: dto.Description,
		Price:       dto.Price,
		Stock: entity.Stock{
			Count: dto.Count,
		},
	}

	return product
}
