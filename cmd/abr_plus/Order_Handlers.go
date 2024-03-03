package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	//"github.com/ddilnaz/shop/pkg/abr-plus/model"
	"github.com/ddilnaz/shop/pkg/abr-plus/model"
	"github.com/gorilla/mux"
	//"github.com/shop/pkg/model"
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
		Title       string `json:"title"`
		Description  string `json:"description"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	order := &model.Order{
		Title:      input.Title,
		Description: input.Description,
	}

	err = app.models.Orders.CreateOrder(order)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
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

    menu, err := app.models.Orders.GetOrderById(id)
    if err != nil {
        app.respondWithError(w, http.StatusNotFound, "404 Not Found")
        return
    }

    app.respondWithJSON(w, http.StatusOK, menu)
}


func (app *application) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    param := vars["orderId"]

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
		Title          *string `json:"title"`
		Description    *string `json:"description"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if input.Title != nil {
		order.Title = *input.Title
	}

	if input.Description != nil {
		order.Description = *input.Description
	}

	
	err = app.models.Orders.UpdateOrder(order)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK,order)
}

func (app *application) deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    param := vars["orderId"]

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
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}