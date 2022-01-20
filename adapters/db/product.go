package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/retatu/arq-hexagonal/application"
)

type ProductDB struct {
	db *sql.DB
}

// func (p *ProductDB) Save(product application.ProductInterface) (application.ProductInterface, error) {
// }

func (p *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
