package usecase

import (
	"fmt"

	"enigmacamp.com/be-enigma-laundry/model"
	"enigmacamp.com/be-enigma-laundry/repository"
)

type CustomerUseCase interface {
	FindById(id string) (model.Customer, error)
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (c *customerUseCase) FindById(id string) (model.Customer, error) {
	customer, err := c.repo.Get(id)
	if err != nil {
		return model.Customer{}, fmt.Errorf("customer with ID %s not found", id)
	}
	return customer, nil
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
