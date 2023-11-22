package repository

import (
	"database/sql"

	"enigmacamp.com/be-enigma-laundry/model"
)

type CustomerRepository interface {
	Get(id string) (model.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) Get(id string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRow(`SELECT id,name,phone_number,address,created_at,updated_at FROM customers WHERE id = $1`, id).
		Scan(
			&customer.Id,
			&customer.Name,
			&customer.PhoneNumber,
			&customer.Address,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)

	if err != nil {
		return model.Customer{}, err
	}

	return customer, nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}
