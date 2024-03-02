package model

import (
	"database/sql"
	"errors"
	"log"
)

var orders = []Order{
	{Id: 1, UserId: 1, ProductItemTitle: "Eiffel Tower Tour", Quantity: 2, TotalPrice: 4000, OrderDate: "2024-03-01", Status: "Pending"},
	{Id: 2, UserId: 2, ProductItemTitle: "Historical Sites Pass", Quantity: 1, TotalPrice: 1800, OrderDate: "2024-03-02", Status: "Confirmed"},
	{Id: 3, UserId: 3, ProductItemTitle: "Island Retreat Package", Quantity: 3, TotalPrice: 4500, OrderDate: "2024-03-03", Status: "Shipped"},
	{Id: 4, UserId: 4, ProductItemTitle: "Mountain Trekking Adventure", Quantity: 1, TotalPrice: 3000, OrderDate: "2024-03-04", Status: "Delivered"},
	{Id: 5, UserId: 5, ProductItemTitle: "City Sightseeing Tour", Quantity: 2, TotalPrice: 4200, OrderDate: "2024-03-05", Status: "Pending"},
}


type OrderModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
func (om *OrderModel) GetAllOrders() ([]Order, error) {
	return orders, nil
}

func (om *OrderModel) GetOrderByID(id int) (*Order, error) {
	for _, o := range orders {
		if o.Id == id {
			return &o, nil
		}
	}
	return nil, errors.New("order not found")
}

func (om *OrderModel) CreateOrder(order *Order) error {
	order.Id = len(orders) + 1
	orders = append(orders, *order)
	return nil
}

func (om *OrderModel) UpdateOrder(order *Order) error {
	for i, o := range orders {
		if o.Id == order.Id {
			orders[i] = *order
			return nil
		}
	}
	return errors.New("order not found")
}

func (om *OrderModel) DeleteOrder(id int) error {
	for i, order := range orders {
		if order.Id == id {
			// Remove order from slice
			orders = append(orders[:i], orders[i+1:]...)
			return nil
		}
	}
	return errors.New("order not found")
}
