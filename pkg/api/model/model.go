package model

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

var users = []User{
	{Name: "Zeinolla Dilnaz", Username: "zeinolla_d", Password: "password123"},
	{Name: "Zhumanova Zhanel", Username: "ahanel_zh", Password: "pass456"},
	{Name: "Ali John", Username: "ali_zh", Password: "secret123"},
	{Name: "Kami White", Username: "kami_w", Password: "exam_pass"},
	{Name: "Airat Green", Username: "airat_g", Password: "green"},
}


var productItems = []ProductItem{
	{Title: "Eiffel Tower Tour", Description: "Guided tour of the iconic landmark. Explore the history and architecture of the Eiffel Tower with our knowledgeable guides.", Price: 2000, Image: "eiffel_tower.jpg"},
	{Title: "Island Retreat Package", Description: "Indulge in an all-inclusive resort stay on a tropical island. Relax on pristine beaches, enjoy gourmet cuisine, and partake in various activities.", Price: 1500, Image: "tropical_resort.jpg"},
	{Title: "Historical Sites Pass", Description: "Gain access to various historical sites and landmarks. Immerse yourself in the rich history of civilizations with our Historical Sites Pass.", Price: 1800, Image: "historical_sites.jpg"},
	{Title: "Mountain Trekking Adventure", Description: "Embark on a thrilling trek in the mountains. Traverse challenging trails and experience breathtaking landscapes on our Mountain Trekking Adventure.", Price: 3000, Image: "mountain_trek.jpg"},
	{Title: "City Sightseeing Tour", Description: "Explore the city's landmarks and attractions. Our City Sightseeing Tour offers a comprehensive experience of the vibrant urban environment.", Price: 2100, Image: "city_tour.jpg"},
}

var orders = []Order{
	{UserId: 1, ProductItemTitle: "Eiffel Tower Tour", Quantity: 2, TotalPrice: 4000, OrderDate: "2024-03-01", Status: "Pending"},
	{UserId: 2, ProductItemTitle: "Historical Sites Pass", Quantity: 1, TotalPrice: 1800, OrderDate: "2024-03-02", Status: "Confirmed"},
	{UserId: 3, ProductItemTitle: "Island Retreat Package", Quantity: 3, TotalPrice: 4500, OrderDate: "2024-03-03", Status: "Shipped"},
	{UserId: 4, ProductItemTitle: "Mountain Trekking Adventure", Quantity: 1, TotalPrice: 3000, OrderDate: "2024-03-04", Status: "Delivered"},
	{UserId: 5, ProductItemTitle: "City Sightseeing Tour", Quantity: 2, TotalPrice: 4200, OrderDate: "2024-03-05", Status: "Pending"},
}