package order

type CreateOrderDTO struct {
	CustomerID string  `json:"customer_id" binding:"required"`
	Total      float64 `json:"total" binding:"required"`
}
