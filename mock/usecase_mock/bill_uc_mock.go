package usecase_mock

import (
	"enigmacamp.com/be-enigma-laundry/model"
	"enigmacamp.com/be-enigma-laundry/model/dto"
	"github.com/stretchr/testify/mock"
)

type BillUseCaseMock struct {
	mock.Mock
}

func (b *BillUseCaseMock) FindById(id string) (model.Bill, error) {
	args := b.Called(id)
	return args.Get(0).(model.Bill), args.Error(1)
}

func (b *BillUseCaseMock) RegisterNewBill(payload dto.BillRequestDto) (model.Bill, error) {
	args := b.Called(payload)
	return args.Get(0).(model.Bill), args.Error(1)
}
