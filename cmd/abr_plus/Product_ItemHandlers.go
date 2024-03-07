//C:\Users\Lenovo\Desktop\shop\cmd\abr_plus\Product_ItemHandlers.go
package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/ddilnaz/shop/pkg/abr-plus/model"
	"strconv"
	
)


func (app *application) CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title          string `json:"title"`
		Description    string `json:"description"`
		Price      int`json:"price"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	product_item := &model.ProductItem{
		Title:           input.Title,
		Description:    		input.Description,
		Price: input.Price,
	}

	err = app.models.ProductItems.CreateProductItem(product_item)
	if err != nil {
        log.Printf("Error creating item: %s\n", err)
        app.respondWithError(w, http.StatusInternalServerError, "Error creating item")
        return
    }

	app.respondWithJSON(w, http.StatusCreated, product_item)
}


func (app *application) deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["user_id"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	err = app.models.ProductItems.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}