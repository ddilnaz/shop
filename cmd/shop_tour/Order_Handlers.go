// C:\Users\Lenovo\Desktop\shop\cmd\abr_plus\Order_Handlers.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/ddilnaz/shop/pkg/shop_tour/model"
	"github.com/gorilla/mux"
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
		UserID      int    `json:"user_id"`
		ItemID      int    `json:"item_id"` 
		Status      string `json:"status"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		log.Print("zdes")
		app.respondWithError(w, http.StatusBadRequest, "Неверный формат запроса")
		return
	}

	order := &model.Order{
		UserID:      input.UserID,
		ItemID:      input.ItemID,
		Status:      input.Status,
	}

	if order.Status == "" {
		order.Status = "Pending"
	}

	err = app.models.Orders.CreateOrder(order, input.UserID)
	if err != nil {
		log.Printf("Ошибка создания заказа: %s\n", err)
		app.respondWithError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, order)
}

func (app *application) getOrderHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    param := vars["id"] 

    id, err := strconv.Atoi(param)
    if err != nil || id < 1 {
        app.respondWithError(w, http.StatusBadRequest, "Invalid order ID")
        return
    }

    menu, err := app.models.Orders.GetOrderById(id)
    if err != nil {
        app.respondWithError(w, http.StatusNotFound, "404 Not Found")
        return
    }

    app.respondWithJSON(w, http.StatusOK, menu)
}


func (app *application) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["id"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	order, err := app.models.Orders.GetOrderById(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		Status      *string `json:"status"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if input.Status != nil {
		order.Status = *input.Status
	}

	err = app.models.Orders.UpdateOrder(order)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, order)
}

func (app *application) deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["id"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	err = app.models.Orders.DeleteOrder(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid JSON payload")
		return err
	}
	return nil
}
