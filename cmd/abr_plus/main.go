package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/ddilnaz/shop/pkg/abr-plus/model"
	_ "github.com/lib/pq"
	"log"
	"net/http"
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

	// Use addr in ListenAndServe function
	addr := fmt.Sprintf(":%s", app.config.port)
	fmt.Printf("Server is listening on %s...\n", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

const insertOrderQuery = `
INSERT INTO orders (user_id, product_item_title, quantity, total_price, order_date, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id
`

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
