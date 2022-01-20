package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/retatu/arq-hexagonal/adapters/db"
	"github.com/stretchr/testify/require"
)

var DB *sql.DB

func setup() {
	DB, _ = sql.Open("sqlite3", ":memory:")
	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	table := `create table products("id" string, "name" string, "price" float, "status" string)`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Teste", 0, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setup()
	defer DB.Close()

	productDB := db.NewProductDB(DB)
	product, err := productDB.Get("abc")
	require.Nil(t, err)
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "Teste", product.GetName())
	require.Equal(t, "disabled", product.GetStatus())
}
