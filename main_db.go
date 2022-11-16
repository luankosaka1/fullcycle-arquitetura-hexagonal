package main

import (
	"database/sql"
	db2 "github.com/luankosaka1/fullcycle-arquitetura-hexagonal/adapters/db"
	"github.com/luankosaka1/fullcycle-arquitetura-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product 1", 10.50)

	productService.Enable(product)
}
