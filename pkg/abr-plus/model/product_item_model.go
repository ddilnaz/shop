//C:\Users\Lenovo\Desktop\shop\pkg\abr-plus\model\product_item_model.go
package model

import (
	"database/sql"
	"log"
	"context"
	"time"
)

type ProductItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	UpdatedAt      string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}
var productItems = []ProductItem{
	{Id: 1, Title: "Eiffel Tower Tour", Description: "Guided tour of the iconic landmark", Price: 2000},
	{Id: 2, Title: "Island Retreat Package", Description: "All-inclusive resort stay", Price: 1500},
	{Id: 3, Title: "Historical Sites Pass", Description: "Access to various historical sites", Price: 1800},
	{Id: 4, Title: "Mountain Trekking Adventure", Description: "Thrilling trek in the mountains", Price: 3000},
	{Id: 5, Title: "City Sightseeing Tour", Description: "Explore the city's landmarks", Price: 2100},
}

type ProductItemModel struct{
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
func GetProductItems() []ProductItem {
	return productItems
}

func (m ProductItemModel) CreateProductItem(product_item *ProductItem) error {
	// Insert a new menu item into the database.
	query := `
		INSERT INTO product_item (title, description, price) 
		VALUES ($1, $2, $3) 
		RETURNING id
		`
	args := []interface{}{product_item.Title, product_item.Description, product_item.Price}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&product_item.Id)
}
	func (m ProductItemModel) UpdateProductItem(product_item *ProductItem) error {
		// Update a specific product_item  in the database.
		query := `
			UPDATE product_item
			SET title = $1, description = $2, price = $3
			WHERE id = $4
			RETURNING updated_at
			`
		args := []interface{}{product_item.Title, product_item.Description, product_item.Price}
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
	
		return m.DB.QueryRowContext(ctx, query, args...).Scan(&product_item.UpdatedAt)
	}


	func (m ProductItemModel) Delete(id int) error {
		query := `
			DELETE FROM product_item
			WHERE id = $1
		`
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
	
		_, err := m.DB.ExecContext(ctx, query, id)
		return err
	}
	
