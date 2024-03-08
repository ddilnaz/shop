//C:\Users\Lenovo\Desktop\shop\cmd\abr_plus\User_Handlers.go
package main

import (
	// "encoding/json"
	// "github.com/ddilnaz/shop/pkg/abr-plus/model"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/ddilnaz/shop/pkg/shop_tour/model"
	"strconv"
)


func (app *application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name          string `json:"name"`
		Email    string `json:"email"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user := &model.User{
		Name:           input.Name,
		Email:    		input.Email,
	}

	err = app.models.Users.CreateUser(user)
	if err != nil {
        log.Printf("Error creating user: %s\n", err)
        app.respondWithError(w, http.StatusInternalServerError, "Error creating user")
        return
    }

	app.respondWithJSON(w, http.StatusCreated, user)
}

func (app *application) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["user_id"]



	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	menu, err := app.models.Users.GetUserById(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, menu)
}

func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["user_id"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	user, err := app.models.Users.GetUserById(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		Name          string `json:"name"`
		Email    string `json:"email"`
	}
	
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if input.Name != "" {
		user.Name = input.Name
	}
	
	if input.Email != "" {
		user.Email = input.Email
	}

	err = app.models.Users.UpdateUser(user)
	if err != nil {
		log.Println("Error updating user:", err)
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}
	
	app.respondWithJSON(w, http.StatusOK, user)
}

func (app *application) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["user_id"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	err = app.models.Users.DeleteUser(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}