//C:\Users\Lenovo\Desktop\shop\pkg\abr-plus\model\items.go
package model

import (
	"database/sql"
	"log"
	"os"
)

type Models struct {
	Users       UserModel
	Orders      OrderModel
	ProductItems ProductItemModel
}

func NewModels(db *sql.DB) Models {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{
		Users: UserModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Orders: OrderModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		ProductItems: ProductItemModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
	}
}