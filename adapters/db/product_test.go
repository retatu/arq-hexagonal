package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func setup() {
	DB, _ sql.Open("sqlite3", ":memory:")
	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB){
	table := `create table products("id" string, "name" string, "price" float, "status" string)`
	stmt, err := db.Prepare(table)
	if err != nil{
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sqlDB){
	insert := `insert into products values("abc", "Teste", 0, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil{
		log.Fatal(err.Error())
	}
	stmt.Exec()
}