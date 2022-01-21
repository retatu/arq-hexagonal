package main

import (
	"database/sql"

	db2 "github.com/retatu/arq-hexagonal/adapters/db"
	"github.com/retatu/arq-hexagonal/application"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDBAdapter := db2.NewProductDB(db)
	productService := application.NewProductService(productDBAdapter)
	product, _ := productService.Create("Example", 30.2)
	productService.Enable(product)
}
