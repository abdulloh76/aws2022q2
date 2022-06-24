package productDto

import "awsshop/internal/entity"

type ProductResponse struct {
	ID          string
	Count       uint64
	Title       string
	Description string
	Price       float32
}

func EntitytoResponsetDTO(product *entity.Product) ProductResponse {
	getProductDto := ProductResponse{
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Count:       product.Stock.Count,
	}

	return getProductDto
}
