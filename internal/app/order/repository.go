package order

import (
	"context"
	"database/sql"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (o *orderRepository) ListOrders(c context.Context) ([]Order, error) {
	rows, err := o.db.Query("SELECT id, customer_id, status, total, created_at FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.Status, &order.Total, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// Create implements OrderRepository.
func (o *orderRepository) Create(c context.Context, order *Order) error {
	panic("unimplemented")

	// collection := tr.database.Collection(tr.collection)

	// _, err := collection.InsertOne(c, task)

	// return err
}

// FetchByCustomerID implements OrderRepository.
func (o *orderRepository) FetchByCustomerID(c context.Context, customerId string) ([]Order, error) {
	panic("unimplemented")

	// collection := tr.database.Collection(tr.collection)

	// var tasks []Task

	// idHex, err := primitive.ObjectIDFromHex(userID)
	// if err != nil {
	// 	return tasks, err
	// }

	// cursor, err := collection.Find(c, bson.M{"userID": idHex})
	// if err != nil {
	// 	return nil, err
	// }

	// err = cursor.All(c, &tasks)
	// if tasks == nil {
	// 	return []Task{}, err
	// }

	// return tasks, err
}
