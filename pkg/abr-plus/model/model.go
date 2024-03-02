//C:\Users\Lenovo\Desktop\shop\pkg\abr-plus\model\items.go
package model

import (
	"database/sql"
	"log"
	"os"
)

type Models struct {
	User       UserModel
	Order      OrderModel
	ProductItem ProductItemModel
}

func NewModels(db *sql.DB) Models {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{
		User: UserModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Order: OrderModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		ProductItem: ProductItemModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
	}
}