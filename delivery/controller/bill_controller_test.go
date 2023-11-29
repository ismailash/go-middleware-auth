package controller

import (
	"bytes"
	"encoding/json"
	middlewaremock "enigmacamp.com/be-enigma-laundry/mock/middleware_mock"
	usecasemock "enigmacamp.com/be-enigma-laundry/mock/usecase_mock"
	"enigmacamp.com/be-enigma-laundry/model"
	"enigmacamp.com/be-enigma-laundry/model/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type BillControllerTestSuite struct {
	suite.Suite
	bum *usecasemock.BillUseCaseMock
	rg  *gin.RouterGroup
	amm *middlewaremock.AuthMiddlewareMock
}

func (suite *BillControllerTestSuite) SetupTest() {
	suite.bum = new(usecasemock.BillUseCaseMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	suite.amm = new(middlewaremock.AuthMiddlewareMock)
}

func TestBillControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BillControllerTestSuite))
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
				Price: 11000,
				Type:  "Pcs",
			},
			Qty:   1,
			Price: 11000,
		},
	},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockPayload = dto.BillRequestDto{
	CustomerId: "1",
	UserId:     "1",
	BillDetails: []model.BillDetail{
		{
			Product: model.Product{Id: "1"},
			Qty:     1,
		},
	},
}

var mockTokenJwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJiZWJhcyIsImV4cCI6MTcwMTA5Njc2MCwiaWF0IjoxNzAxMDkzMTYwLCJ1c2VySWQiOiJmMGEzMWU3Yi1jYmU0LTQ5MDEtYmRjNy0xYWIxYThhNmJhMDgiLCJyb2xlIjoiZW1wbG95ZWUiLCJzZXJ2aWNlcyI6bnVsbH0.UMEiIJ_OG912wEDG9Vgv1BDL0319P6PNeT2MQbqT690"

func (suite *BillControllerTestSuite) TestCreateHandler_Success() {
	// EKSPEKTASI
	suite.bum.On("RegisterNewBill", mockPayload).Return(dummyBill, nil)
	billController := NewBillController(suite.bum, suite.rg, suite.amm)
	billController.Route()
	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah payload dalam bentuk JSON
	mockPayloadJSON, err := json.Marshal(mockPayload)
	assert.NoError(suite.T(), err)

	// EKSEKUSI
	// Simulasi membuat request ke path /api/v1/bills
	// Authorization
	req, err := http.NewRequest(http.MethodPost, "/api/v1/bills", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwt)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("user", dummyBill.User.Id)
	billController.createHandler(ctx)

	// ASSERTION
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}

func (suite *BillControllerTestSuite) TestCreateHandler_BindingFailed() {
	billController := NewBillController(suite.bum, suite.rg, suite.amm)
	billController.Route()

	// EKSPEKTASI
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/bills", nil)
	assert.NoError(suite.T(), err)

	// EKSEKUSI
	req.Header.Set("Authorization", "Bearer "+mockTokenJwt)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("user", dummyBill.User.Id)
	billController.createHandler(ctx)

	// ASSERTION
	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
}
