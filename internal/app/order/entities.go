package order

import "context"

type Order struct {
	ID            int         `json:"id"`
	CustomerID    string      `json:"customer_id"`
	Status        string      `json:"status"`
	Total         float64     `json:"total"`
	CreatedAt     string      `json:"created_at"`
	PaymentMethod string      `json:"payment_method"`
	Items         []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type OrderEvent struct {
	EventType string `json:"event_type"`
	Order     Order  `json:"order"`
}

// OrderRepository defines the interface for order data operations
type OrderRepository interface {
	Create(c context.Context, order *Order) error
	FetchByCustomerID(c context.Context, customerId string) ([]Order, error)
	ListOrders(c context.Context) ([]Order, error)
}

// OrderUseCase defines the interface for order business logic
type OrderUseCase interface {
	Create(c context.Context, order *Order) error
	FetchByCustomerID(c context.Context, customerId string) ([]Order, error)
	ListOrders(c context.Context) ([]Order, error)
}
