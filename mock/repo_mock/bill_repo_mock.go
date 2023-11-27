package repo_mock

import (
	"enigmacamp.com/be-enigma-laundry/model"
	"github.com/stretchr/testify/mock"
)

type BillRepoMock struct {
	mock.Mock
}

func (b *BillRepoMock) Create(payload model.Bill) (model.Bill, error) {
	args := b.Called(payload)
	return args.Get(0).(model.Bill), args.Error(1)
}

func (b *BillRepoMock) Get(id string) (model.Bill, error) {
	args := b.Called(id)
	return args.Get(0).(model.Bill), args.Error(1)
}
