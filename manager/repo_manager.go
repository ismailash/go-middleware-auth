package manager

import "enigmacamp.com/be-enigma-laundry/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	CustomerRepo() repository.CustomerRepository
	ProductRepo() repository.ProductRepository
	BillRepo() repository.BillRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infra.Conn())
}

func (r *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Conn())
}

func (r *repoManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
