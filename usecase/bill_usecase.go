package usecase

import (
	"fmt"

	"enigmacamp.com/be-enigma-laundry/model"
	"enigmacamp.com/be-enigma-laundry/model/dto"
	"enigmacamp.com/be-enigma-laundry/repository"
)

type BillUseCase interface {
	RegisterNewBill(payload dto.BillRequestDto) (model.Bill, error)
	FindById(id string) (model.Bill, error)
}

type billUseCase struct {
	repo       repository.BillRepository
	userUC     UserUseCase
	customerUC CustomerUseCase
	productUC  ProductUseCase
}

func (b *billUseCase) FindById(id string) (model.Bill, error) {
	bill, err := b.repo.Get(id)
	if err != nil {
		return model.Bill{}, fmt.Errorf("bill with ID %s not found", id)
	}
	return bill, nil
}

func (b *billUseCase) RegisterNewBill(payload dto.BillRequestDto) (model.Bill, error) {
	customer, err := b.customerUC.FindById(payload.CustomerId)
	if err != nil {
		return model.Bill{}, err
	}

	user, err := b.userUC.FindById(payload.UserId)
	if err != nil {
		return model.Bill{}, err
	}

	var billDetails []model.BillDetail
	for _, v := range payload.BillDetails {
		product, err := b.productUC.FindById(v.Product.Id)
		if err != nil {
			return model.Bill{}, err
		}

		billDetails = append(billDetails, model.BillDetail{Product: product, Qty: v.Qty, Price: product.Price})
	}

	newBillPayload := model.Bill{
		Customer:    customer,
		User:        user,
		BillDetails: billDetails,
	}

	bill, err := b.repo.Create(newBillPayload)
	if err != nil {
		return model.Bill{}, err
	}

	return bill, nil
}

func NewBillUseCase(
	repo repository.BillRepository,
	userUC UserUseCase,
	customerUC CustomerUseCase,
	productUC ProductUseCase,
) BillUseCase {
	return &billUseCase{
		repo:       repo,
		userUC:     userUC,
		customerUC: customerUC,
		productUC:  productUC,
	}
}
