package usecase

import (
	repomock "enigmacamp.com/be-enigma-laundry/mock/repo_mock"
	usecasemock "enigmacamp.com/be-enigma-laundry/mock/usecase_mock"
	"enigmacamp.com/be-enigma-laundry/model"
	"enigmacamp.com/be-enigma-laundry/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type BillUseCaseTestSuite struct {
	suite.Suite
	brm *repomock.BillRepoMock
	uum *usecasemock.UserUseCaseMock
	cum *usecasemock.CustomerUseCaseMock
	pum *usecasemock.ProductUseCaseMock
	bu  BillUseCase
}

func (suite *BillUseCaseTestSuite) SetupTest() {
	suite.brm = new(repomock.BillRepoMock)
	suite.uum = new(usecasemock.UserUseCaseMock)
	suite.cum = new(usecasemock.CustomerUseCaseMock)
	suite.pum = new(usecasemock.ProductUseCaseMock)
	suite.bu = NewBillUseCase(suite.brm, suite.uum, suite.cum, suite.pum)
}

func TestBillUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BillUseCaseTestSuite))
}

var dummyBill = model.Bill{
	Id:       "1",
	BillDate: time.Now(),
	Customer: model.Customer{
		Id:   "1",
		Name: "Jojo",
	},
	User: model.User{
		Id:   "1",
		Name: "Shinta",
	},
	BillDetails: []model.BillDetail{
		{
			Id:     "1",
			BillId: "1",
			Product: model.Product{
				Id:    "1",
				Name:  "Cuci + Setrika",
				Price: 10000,
				Type:  "Kg",
			},
			Qty:   1,
			Price: 10000,
		},
	},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockPayload = dto.BillRequestDto{
	CustomerId: "1",
	UserId:     "1",
	BillDetails: []model.BillDetail{{
		Product: model.Product{Id: "1"},
		Qty:     1,
	}},
}

// TEST CASE
func (suite *BillUseCaseTestSuite) TestRegisterNewBill_Success() {
	// EKSPEKTASI
	suite.cum.On("FindById", mockPayload.CustomerId).Return(dummyBill.Customer, nil)

	suite.uum.On("FindById", mockPayload.UserId).Return(dummyBill.User, nil)

	var mockBillDetails []model.BillDetail
	for _, v := range mockPayload.BillDetails {
		suite.pum.On("FindById", v.Product.Id).Return(dummyBill.BillDetails[0].Product, nil)

		mockBillDetails = append(mockBillDetails, model.BillDetail{Product: dummyBill.BillDetails[0].Product, Qty: v.Qty, Price: dummyBill.BillDetails[0].Product.Price})
	}

	mockNewBillPayload := model.Bill{
		Customer:    dummyBill.Customer,
		User:        dummyBill.User,
		BillDetails: mockBillDetails,
	}

	// EKSEKUSI
	suite.brm.On("Create", mockNewBillPayload).Return(dummyBill, nil)
	_, err := suite.bu.RegisterNewBill(mockPayload)

	// ASSERT
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}
