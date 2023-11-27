package usecase_mock

import (
	"enigmacamp.com/be-enigma-laundry/model"
	"github.com/stretchr/testify/mock"
)

type CustomerUseCaseMock struct {
	mock.Mock
}

func (c *CustomerUseCaseMock) FindById(id string) (model.Customer, error) {
	args := c.Called(id)
	return args.Get(0).(model.Customer), args.Error(1)
}
