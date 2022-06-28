package productDto

import "awsshop/internal/entity"

type ProductResponse struct {
	ID          string  `json:"id"`
	Count       uint64  `json:"count"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func EntitytoResponsetDTO(product *entity.Product) ProductResponse {
	getProductDto := ProductResponse{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Count:       product.Stock.Count,
	}

	return getProductDto
}
