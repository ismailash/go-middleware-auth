package usecase_mock

import (
	"enigmacamp.com/be-enigma-laundry/model"
	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (u *UserUseCaseMock) FindById(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserUseCaseMock) FindByUsernamePassword(username string, password string) (model.User, error) {
	args := u.Called(username, password)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserUseCaseMock) RegisterNewUser(payload model.User) (model.User, error) {
	args := u.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}
