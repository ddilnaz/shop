//C:\Users\Lenovo\Desktop\shop\cmd\abr_plus\main.go
package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/ddilnaz/shop/pkg/abr-plus/model"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":8081", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:postgres@localhost/shop?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	app := &application{
		config: cfg,
		models: model.NewModels(db),
	}

	app.run()
}

func (app *application) run() {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Menu Singleton
	
	r.HandleFunc("/orders/{id:[0-9]+}", app.updateOrderHandler).Methods("PUT")
	v1.HandleFunc("/orders/{id:[0-9]+}", app.getOrderHandler).Methods("GET")
	v1.HandleFunc("/orders/{id:[0-9]+}", app.deleteOrderHandler).Methods("DELETE") 
	v1.HandleFunc("/orders", app.createOrderHandler).Methods("POST")

	log.Printf("Starting server on %s\n", app.config.port)
	err := http.ListenAndServe(app.config.port, r)
	log.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
