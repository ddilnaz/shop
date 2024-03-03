package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
    "github.com/shop/pkg/model"
	"net/http"
	"strconv"
)

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *application) createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID           int    `json:"user_id"`
		ProductItemTitle string `json:"product_item_title"`
		Quantity         int    `json:"quantity"`
		TotalPrice       int    `json:"total_price"`
		OrderDate        string `json:"order_date"`
		Status           string `json:"status"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	order := &model.Order{
		// Здесь нужно сконструировать объект заказа в соответствии с вашей логикой.
		// Пример: Id: 0, UserID: input.UserID, ProductItemTitle: input.ProductItemTitle, и так далее.
	}

	err = app.models.Order.CreateOrder(order)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Failed to create order")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, order)
}

func (app *application) getOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["orderId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	order, err := app.models.Order.GetOrderById(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "Order not found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, order)
}
func (app *application) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["orderId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	order, err := app.models.Order.GetOrderById(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "Order not found")
		return
	}

	var input struct {
		UserID           *int    `json:"user_id"`
		ProductItemTitle *string `json:"product_item_title"`
		Quantity         *int    `json:"quantity"`
		TotalPrice       *int    `json:"total_price"`
		OrderDate        *string `json:"order_date"`
		Status           *string `json:"status"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if input.UserID != nil {
		order.UserID = *input.UserID
	}

	if input.ProductItemTitle != nil {
		order.ProductItemTitle = *input.ProductItemTitle
	}

	if input.Quantity != nil {
		order.Quantity = *input.Quantity
	}

	if input.TotalPrice != nil {
		order.TotalPrice = *input.TotalPrice
	}

	if input.OrderDate != nil {
		order.OrderDate = *input.OrderDate
	}

	if input.Status != nil {
		order.Status = *input.Status
	}

	err = app.models.Order.UpdateOrder(order)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Failed to update order")
		return
	}

	app.respondWithJSON(w, http.StatusOK, order)
}

func (app *application) deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["orderId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	err = app.models.Order.DeleteOrder(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Failed to delete order")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
