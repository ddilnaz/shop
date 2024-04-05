// C:\Users\Lenovo\Desktop\shop\pkg\shop_tour\model\order_model.go
package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Order struct {
	OrderID     int    `json:"order_id" db:"order_id"`
	UserID      int    `json:"user_id" db:"user_id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Status      string `json:"status"`
	ItemID      int    `json:"item_id" db:"item_id"` 
}

type OrderModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m *OrderModel) CreateOrder(order *Order, userID int) error {
	query := `
		INSERT INTO orders (status, user_id, item_id) 
		VALUES ($1, $2, $3) 
		RETURNING order_id, created_at, updated_at, status 
	`

	args := []interface{}{order.Status, order.UserID, order.ItemID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&order.OrderID, &order.CreatedAt, &order.UpdatedAt, &order.Status)
}


func (m OrderModel) GetOrderById(orderID int) (*Order, error) {
	query := `
		SELECT order_id, user_id, created_at, updated_at, status 
		FROM orders 
		WHERE order_id = $1
	`

	var order Order
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, orderID)
	err := row.Scan(&order.OrderID, &order.UserID, &order.CreatedAt, &order.UpdatedAt, &order.Status)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (m *OrderModel) UpdateOrder(order *Order) error {
	query := `
		UPDATE orders SET  status = $1
		WHERE order_id = $2
		RETURNING updated_at
	`
	args := []interface{}{order.Status, order.OrderID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&order.UpdatedAt)
}

func (m OrderModel) DeleteOrder(orderID int) error {
	query := `
		DELETE FROM orders WHERE order_id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, orderID)
	return err
}
