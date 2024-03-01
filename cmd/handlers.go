package main

import (
	"net/http"
	"strconv"
	//"pkg\api\model\model.go"
	"github.com/gorilla/mux"
)
func (app *application) createTourHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Price       int    `json:"price"`
		Image       string `json:"image"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	tour := &model.ProductItem{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Image:       input.Image,
	}

	err = app.models.Tours.Insert(tour)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, tour)
}

func (app *application) getTourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["tourId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid tour ID")
		return
	}

	tour, err := app.models.Tours.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, tour)
}

func (app *application) updateTourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["tourId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid tour ID")
		return
	}

	tour, err := app.models.Tours.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Price       *int    `json:"price"`
		Image       *string `json:"image"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if input.Title != nil {
		tour.Title = *input.Title
	}

	if input.Description != nil {
		tour.Description = *input.Description
	}

	if input.Price != nil {
		tour.Price = *input.Price
	}

	if input.Image != nil {
		tour.Image = *input.Image
	}

	err = app.models.Tours.Update(tour)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, tour)
}

func (app *application) deleteTourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["tourId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid tour ID")
		return
	}

	err = app.models.Tours.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
