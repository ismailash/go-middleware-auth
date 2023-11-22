package usecase

import (
	"errors"
	"fmt"

	"enigmacamp.com/be-enigma-laundry/model"
	"enigmacamp.com/be-enigma-laundry/repository"
	"enigmacamp.com/be-enigma-laundry/utils/common"
)

type UserUseCase interface {
	FindById(id string) (model.User, error)
	FindByUsernamePassword(username string, password string) (model.User, error)
	RegisterNewUser(payload model.User) (model.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) FindByUsernamePassword(username string, password string) (model.User, error) {
	user, err := u.repo.GetByUsername(username)
	if err != nil {
		return model.User{}, errors.New("invalid username or password")
	}

	// compare password
	if err := common.ComparePasswordHash(user.Password, password); err != nil {
		return model.User{}, err
	}

	// kita set password kosong agar tidak ditampilkan di response
	user.Password = ""
	return user, nil
}

func (u *userUseCase) RegisterNewUser(payload model.User) (model.User, error) {
	// biasanya ada pengecekan jika email / username sudah ada di database, maka tidak boleh diinsert lagi

	// cek role
	if !payload.IsValidRole() {
		return model.User{}, errors.New("invalid role, role must be admin or employee")
	}

	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return model.User{}, err
	}

	payload.Password = newPassword
	return u.repo.Create(payload)
}

func (u *userUseCase) FindById(id string) (model.User, error) {
	user, err := u.repo.Get(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
