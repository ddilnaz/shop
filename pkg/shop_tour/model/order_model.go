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
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type OrderModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m *OrderModel) CreateOrder(order *Order, userID int) error {
	query := `
		INSERT INTO orders ( title, description , status, user_id) 
		VALUES ($1, $2, $3, $4) 
		RETURNING order_id, created_at, updated_at, status
	`

	args := []interface{}{order.Title, order.Description, order.Status, order.UserID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&order.OrderID, &order.CreatedAt, &order.UpdatedAt, &order.Status)
}

func (m OrderModel) GetOrderById(orderID int) (*Order, error) {
	query := `
		SELECT order_id, user_id, created_at, updated_at, title, description, status 
		FROM orders 
		WHERE order_id = $1
	`

	var order Order
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, orderID)
	err := row.Scan(&order.OrderID, &order.UserID, &order.CreatedAt, &order.UpdatedAt, &order.Title, &order.Description, &order.Status)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (m *OrderModel) UpdateOrder(order *Order) error {
	query := `
		UPDATE orders SET title = $1, description = $2, status = $3
		WHERE order_id = $4
		RETURNING updated_at
	`
	args := []interface{}{order.Title, order.Description, order.Status, order.OrderID}
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
