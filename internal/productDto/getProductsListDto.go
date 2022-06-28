package productDto

type ProductsListResponse struct {
	ID          string  `json:"id"`
	Count       uint64  `json:"count"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
