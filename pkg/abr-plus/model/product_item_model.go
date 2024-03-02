package model

import (
	"database/sql"
	"errors"
	"log"
)
var productItems = []ProductItem{
	{Id: 1, Title: "Eiffel Tower Tour", Description: "Guided tour of the iconic landmark", Price: 2000, Image: "eiffel_tower.jpg"},
	{Id: 2, Title: "Island Retreat Package", Description: "All-inclusive resort stay", Price: 1500, Image: "tropical_resort.jpg"},
	{Id: 3, Title: "Historical Sites Pass", Description: "Access to various historical sites", Price: 1800, Image: "historical_sites.jpg"},
	{Id: 4, Title: "Mountain Trekking Adventure", Description: "Thrilling trek in the mountains", Price: 3000, Image: "mountain_trek.jpg"},
	{Id: 5, Title: "City Sightseeing Tour", Description: "Explore the city's landmarks", Price: 2100, Image: "city_tour.jpg"},
}

type ProductItemModel struct{
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}


func (pim *ProductItemModel) GetAllProductItems() ([]ProductItem, error) {
	return productItems, nil
}

func (pim *ProductItemModel) GetProductItemByID(id int) (*ProductItem, error) {
	for _, item := range productItems {
		if item.Id == id {
			return &item, nil
		}
	}
	return nil, errors.New("product item not found")
}

func (pim *ProductItemModel) CreateProductItem(productItem *ProductItem) error {
	// Simulating auto-increment ID
	productItem.Id = len(productItems) + 1
	productItems = append(productItems, *productItem)
	return nil
}

func (pim *ProductItemModel) UpdateProductItem(productItem *ProductItem) error {
	for i, item := range productItems {
		if item.Id == productItem.Id {
			productItems[i] = *productItem
			return nil
		}
	}
	return errors.New("product item not found")
}

func (pim *ProductItemModel) DeleteProductItem(id int) error {
	for i, item := range productItems {
		if item.Id == id {
			// Remove product item from slice
			productItems = append(productItems[:i], productItems[i+1:]...)
			return nil
		}
	}
	return errors.New("product item not found")
}
