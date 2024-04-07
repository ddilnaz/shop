package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// routes is our main application's router.
func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	// Convert the app.notFoundResponse helper to a http.Handler using the http.HandlerFunc()
	// adapter, and then set it as the custom error handler for 404 Not Found responses.
	r.NotFoundHandler = http.HandlerFunc(app.notFoundResponse)

	// Convert app.methodNotAllowedResponse helper to a http.Handler and set it as the custom
	// error handler for 405 Method Not Allowed responses
	r.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowedResponse)

	// healthcheck
	r.HandleFunc("/api/v1/healthcheck", app.healthcheckHandler).Methods("GET")

	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/user/{user_id:[0-9]+}", app.getUsersHandler).Methods("GET")
	v1.HandleFunc("/user/{user_id:[0-9]+}", app.deleteUserHandler).Methods("DELETE")
	v1.HandleFunc("/user/{user_id:[0-9]+}", app.updateUserHandler).Methods("PUT")
	v1.HandleFunc("/user",app.CreateUserHandler).Methods("POST")

	v1.HandleFunc("/items",app.CreateItemHandler).Methods("POST")
//	v1.HandleFunc("/items/{items_id:[0-9]+}",app.deleteItemHandler).Methods("DELETE")
	v1.HandleFunc("/items/{items_id:[0-9]+}" , app.getItemHandler).Methods("GET")
	v1.HandleFunc("/items/{items_id:[0-9]+}", app.updateItemHandler).Methods("PUT")


	v1.HandleFunc("/orders", app.createOrderHandler).Methods("POST")
	v1.HandleFunc("/orders/{id:[0-9]+}", app.updateOrderHandler).Methods("PUT")
	v1.HandleFunc("/orders/{id:[0-9]+}", app.getOrderHandler).Methods("GET")
	v1.HandleFunc("/orders/{id:[0-9]+}", app.deleteOrderHandler).Methods("DELETE") 
	return app.authenticate(r)
}