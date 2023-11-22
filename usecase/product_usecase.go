package usecase

import (
	"fmt"

	"enigmacamp.com/be-enigma-laundry/model"
	"enigmacamp.com/be-enigma-laundry/repository"
)

type ProductUseCase interface {
	FindById(id string) (model.Product, error)
}

type productUseCase struct {
	repo repository.ProductRepository
}

func (p *productUseCase) FindById(id string) (model.Product, error) {
	product, err := p.repo.Get(id)
	if err != nil {
		return model.Product{}, fmt.Errorf("product with ID %s not found", id)
	}
	return product, nil
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}
