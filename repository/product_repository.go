package repository

import (
	"database/sql"

	"enigmacamp.com/be-enigma-laundry/model"
)

type ProductRepository interface {
	Get(id string) (model.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func (p *productRepository) Get(id string) (model.Product, error) {
	var product model.Product
	err := p.db.QueryRow(`SELECT id,name,price,type,created_at,updated_at FROM products WHERE id = $1`, id).
		Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Type,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
