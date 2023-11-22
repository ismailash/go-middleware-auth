package manager

import "enigmacamp.com/be-enigma-laundry/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	CustomerUseCase() usecase.CustomerUseCase
	ProductUseCase() usecase.ProductUseCase
	BillUseCase() usecase.BillUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) BillUseCase() usecase.BillUseCase {
	return usecase.NewBillUseCase(u.repo.BillRepo(), u.UserUseCase(), u.CustomerUseCase(), u.ProductUseCase())
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repo.CustomerRepo())
}

func (u *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repo.ProductRepo())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
