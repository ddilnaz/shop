//C:\Users\Lenovo\Desktop\shop\cmd\abr_plus\Product_ItemHandlers.go
package main

import (
	"net/http"
	"log"
	"github.com/ddilnaz/shop/pkg/shop_tour/model"
	"github.com/ddilnaz/shop/pkg/shop_tour/validator"
	"errors"
	
)


func (app *application) CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title          string `json:"title"`
		Description    string `json:"description"`
		Price          int`json:"price"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	product_item := &model.ProductItem{
		Title:           input.Title,
		Description:     input.Description,
		Price:			 input.Price,
	}

	err = app.models.ProductItems.CreateProductItem(product_item)
	if err != nil {
        log.Printf("Error creating item: %s\n", err)
        app.respondWithError(w, http.StatusInternalServerError, "Error creating item")
        return
    }

	app.writeJSON(w, http.StatusCreated, envelope{"product_item": product_item}, nil)
}

func (app *application) getItemsList(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title              string
		NutritionValueFrom int
		NutritionValueTo   int
		model.Filters
	}
	v := validator.New()
	qs := r.URL.Query()

	// Use our helpers to extract the title and nutrition value range query string values, falling back to the
	// defaults of an empty string and an empty slice, respectively, if they are not provided
	// by the client.
	input.Title = app.readStrings(qs, "title", "")
	input.NutritionValueFrom = app.readInt(qs, "nutritionFrom", 0, v)
	input.NutritionValueTo = app.readInt(qs, "nutritionTo", 0, v)

	// Ge the page and page_size query string value as integers. Notice that we set the default
	// page value to 1 and default page_size to 20, and that we pass the validator instance
	// as the final argument.
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	// Extract the sort query string value, falling back to "id" if it is not provided
	// by the client (which will imply an ascending sort on menu ID).
	input.Filters.Sort = app.readStrings(qs, "sort", "id")

	// Add the supported sort value for this endpoint to the sort safelist.
	// name of the column in the database.
	input.Filters.SortSafeList = []string{
		// ascending sort values
		"id", "title", "nutrition_value",
		// descending sort values
		"-id", "-title", "-nutrition_value",
	}

	if model.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	menus, metadata, err := app.models.ProductItems.GetAll(input.Title, input.NutritionValueFrom, input.NutritionValueTo, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"menus": menus, "metadata": metadata}, nil)
}



func (app *application) getItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	menu, err := app.models.ProductItems.GetItemById(id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"menu": menu}, nil)
}


func (app *application) updateItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	menu, err := app.models.ProductItems.GetItemById(id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title          *string `json:"title"`
		Description    *string `json:"description"`
		Price *uint   `json:"price"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		menu.Title = *input.Title
	}

	if input.Description != nil {
		menu.Description = *input.Description
	}

	// if input.Price != 0 {
	// 	menu.Price = input.Price
	// }
	v := validator.New()

	if model.ValidateMenu(v, menu); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.ProductItems.UpdateProductItem(menu)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"menu": menu}, nil)
}

func (app *application) deleteMenuHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.ProductItems.DeleteById(id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "success"}, nil)
}
// func (app *application) updateItemHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	param := vars["items_id"]

// 	id, err := strconv.Atoi(param)
// 	if err != nil || id < 1 {
// 		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
// 		return
// 	}

// 	item, err := app.models.ProductItems.GetItemById(id)
// 	if err != nil {
// 		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
// 		return
// 	}

// 	var input struct {
// 		Title          string `json:"title"`
// 		Description    string `json:"description"`
// 		Price          int `json:"price"`
// 	}
	
// 	err = app.readJSON(w, r, &input)
// 	if err != nil {
// 		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	if input.Title != "" {
// 		item.Title = input.Title
// 	}
	
// 	if input.Description != "" {
// 		item.Description = input.Description
// 	}
// 	if input.Price != 0 {
// 		item.Price = input.Price
// 	}	
// 	err = app.models.ProductItems.UpdateProductItem(item)
// 	if err != nil {
// 		log.Println("Error updating user:", err)
// 		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
// 		return
// 	}
	
// 	app.respondWithJSON(w, http.StatusOK, item)
// }
// func (app *application) deleteItemHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	param := vars["items_id"]

// 	id, err := strconv.Atoi(param)
// 	if err != nil || id < 1 {
// 		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
// 		return
// 	}

// 	err = app.models.ProductItems.DeleteById(id)
// 	if err != nil {
// 		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
// 		return
// 	}

// 	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
// }