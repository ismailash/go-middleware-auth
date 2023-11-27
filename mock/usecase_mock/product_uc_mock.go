package usecase_mock

import (
	"enigmacamp.com/be-enigma-laundry/model"
	"github.com/stretchr/testify/mock"
)

type ProductUseCaseMock struct {
	mock.Mock
}

func (c *ProductUseCaseMock) FindById(id string) (model.Product, error) {
	args := c.Called(id)
	return args.Get(0).(model.Product), args.Error(1)
}
