//C:\Users\Lenovo\Desktop\shop\pkg\abr-plus\model\items.go
package model
import (
 	//"errors"
)

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


type ProductItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	Image       string `json:"image" db:"image"`
}

type Order struct {
	Id             int    `json:"id" db:"id"`
	UserId         int    `json:"user_id" db:"user_id"`
	ProductItemTitle string `json:"product_item_title" db:"product_item_title"`
	Quantity       int    `json:"quantity" db:"quantity"`
	TotalPrice     int    `json:"total_price" db:"total_price"`
	OrderDate      string `json:"order_date" db:"order_date"`
	Status         string `json:"status" db:"status"`
}
