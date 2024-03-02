//cmd\abr_plus\handlers.go
package main

import (
	"encoding/json"

	"github.com/ddilnaz/shop/pkg/abr-plus/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)


func (app *application) createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
	Id             int    `json:"id" db:"id"`
	UserId         int    `json:"user_id" db:"user_id"`
	ProductItemTitle string `json:"product_item_title" db:"product_item_title"`
	Quantity       int    `json:"quantity" db:"quantity"`
	TotalPrice     int    `json:"total_price" db:"total_price"`
	OrderDate      string `json:"order_date" db:"order_date"`
	Status         string `json:"status" db:"status"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	order := &model.Order{
		Id:          input.Id,
		UserId:    input.UserId,
		ProductItemTitle: input.ProductItemTitle,
		Quantity: input.Quantity,
		TotalPrice: input.TotalPrice,
		OrderDate: input.OrderDate,
		Status: input.Status,
	}

	err = app.models.Order.CreateOrder(order)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, order)
}


func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(data); err != nil {
		return err
	}

	return nil
}

func (app *application) respondWithError(w http.ResponseWriter, status int, message string) {
	response := map[string]string{"error": message}
	app.respondWithJSON(w, status, response)
}

func (app *application) respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
// func (app *application) createOrderHandler(w http.ResponseWriter, r *http.Request) {
// 	var input model.Order
// 	err := app.readJSON(w, r, &input)
// 	if err != nil {
// 		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	err = app.models.Order.CreateOrder(&input)
// 	if err != nil {
// 		app.respondWithError(w, http.StatusInternalServerError, "Failed to create order")
// 		return
// 	}

// 	app.respondWithJSON(w, http.StatusCreated, input)
// }

func (app *application) getTourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid tour ID")
		return
	}
	tour, err := app.models.ProductItem.GetProductItemByID(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "Tour not found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, tour)
}


// func (app *application) updateTourHandler(w http.ResponseWriter, r *http.Request) {
	
// }

// func (app *application) deleteTourHandler(w http.ResponseWriter, r *http.Request) {
	
// }

