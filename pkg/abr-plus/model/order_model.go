//C:\Users\Lenovo\Desktop\shop\pkg\abr-plus\model\order_model.go
package model

import (
	"database/sql"
	"log"
	"time"
	"context"
)

type Order struct {
	Id             string `json:"id"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Status         string `json:"status" db:"status"`
}

type OrderModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m *OrderModel) CreateOrder(order *Order) error {
	query := `
		INSERT INTO "order" (id, title, description, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, status, created_at, updated_at
	`
		
	args := []interface{}{order.Id, order.Title, order.Description}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&order.Status, &order.CreatedAt, &order.UpdatedAt)
}
	


func (m OrderModel) GetOrderById(id int) (*Order, error) {
	// Retrieve a specific menu item based on its ID.
	query := `
		SELECT id, created_at, updated_at, title, description, status
		FROM menu
		WHERE id = $1
		`
	var order Order
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&order.Id, &order.CreatedAt, &order.UpdatedAt, &order.Title, &order.Description, &order.Status)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
func (m *OrderModel) UpdateOrder(order *Order) error {
	query := `
		UPDATE "order"
		SET title = $1, description = $2, status = $3
		WHERE id = $4
		RETURNING updated_at
	`
	args := []interface{}{order.Title, order.Description, order.Status, order.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&order.UpdatedAt)
}

func (m OrderModel) DeleteOrder(id int) error {
	// Delete a specific menu item from the database.
	query := `
		DELETE FROM menu
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}