//C:\Users\Lenovo\Desktop\shop\pkg\abr-plus\model\product_item_model.go
package model

import (
	"database/sql"
	"log"
	"context"
	"time"
	"github.com/ddilnaz/shop/pkg/shop_tour/validator"
	"fmt"
)

type ProductItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	UpdatedAt      string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}
// var productItems = []ProductItem{
// 	{Id: 1, Title: "Eiffel Tower Tour", Description: "Guided tour of the iconic landmark", Price: 2000},
// 	{Id: 2, Title: "Island Retreat Package", Description: "All-inclusive resort stay", Price: 1500},
// 	{Id: 3, Title: "Historical Sites Pass", Description: "Access to various historical sites", Price: 1800},
// 	{Id: 4, Title: "Mountain Trekking Adventure", Description: "Thrilling trek in the mountains", Price: 3000},
// 	{Id: 5, Title: "City Sightseeing Tour", Description: "Explore the city's landmarks", Price: 2100},
// }

type ProductItemModel struct{
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m ProductItemModel) GetAll(title string, from, to int, filters Filters) ([]*ProductItem, Metadata, error) {

	// Retrieve all menu items from the database.
	query := fmt.Sprintf(
		`
		SELECT count(*) OVER(), id, created_at, updated_at, title, description, price
		FROM product_item
		WHERE (LOWER(title) = LOWER($1) OR $1 = '')
		AND (price >= $2 OR $2 = 0)
		AND (price <= $3 OR $3 = 0)
		ORDER BY %s %s, id ASC
		LIMIT $4 OFFSET $5
		`,
		filters.sortColumn(), filters.sortDirection())

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Organize our four placeholder parameter values in a slice.
	args := []interface{}{title, from, to, filters.limit(), filters.offset()}

	// log.Println(query, title, from, to, filters.limit(), filters.offset())
	// Use QueryContext to execute the query. This returns a sql.Rows result set containing
	// the result.
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	// Importantly, defer a call to rows.Close() to ensure that the result set is closed
	// before GetAll returns.
	defer func() {
		if err := rows.Close(); err != nil {
			m.ErrorLog.Println(err)
		}
	}()

	// Declare a totalRecords variable
	totalRecords := 0

	var items []*ProductItem
	for rows.Next() {
		var item ProductItem
		err := rows.Scan(&totalRecords, &item.Id, &item.CreatedAt, &item.UpdatedAt, &item.Title, &item.Description, &item.Price)
		if err != nil {
			return nil, Metadata{}, err
		}

		// Add the Movie struct to the slice
		items = append(items, &item)
	}

	// When the rows.Next() loop has finished, call rows.Err() to retrieve any error
	// that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	// Generate a Metadata struct, passing in the total record count and pagination parameters
	// from the client.
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	// If everything went OK, then return the slice of the movies and metadata.
	return items, metadata, nil
}
func (m ProductItemModel) GetItemById(id int) (*ProductItem, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Retrieve a specific order item based on its ID.
	query := `
		SELECT id, created_at, updated_at, title, description, price 
		FROM product_item
		WHERE id = $1
		`
	var items ProductItem
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&items.Id, &items.CreatedAt, &items.UpdatedAt, &items.Title, &items.Description, &items.Price)
	if err != nil {
		return nil, err
	}
	return &items, nil
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
	args := []interface{}{product_item.Title, product_item.Description, product_item.Price,product_item.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&product_item.UpdatedAt)
}

func (m ProductItemModel) DeleteById(id int) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM product_item
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}
	

func ValidateMenu(v *validator.Validator, item *ProductItem) {
	// Check if the title field is empty.
	v.Check(item.Title != "", "title", "must be provided")
	// Check if the title field is not more than 100 characters.
	v.Check(len(item.Title) <= 100, "title", "must not be more than 100 bytes long")
	// Check if the description field is not more than 1000 characters.
	v.Check(len(item.Description) <= 1000, "description", "must not be more than 1000 bytes long")
	// Check if the nutrition value is not more than 10000.
	v.Check(item.Price > 0, "price", "must not be more than 0")
}