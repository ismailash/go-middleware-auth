package repository

import (
	"database/sql"
	"enigmacamp.com/be-enigma-laundry/model"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

// Buat suite
type BillRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    BillRepository
}

// Setup
func (suite *BillRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.sqlmock = sqlmock
	suite.repo = NewBillRepository(suite.mockDB)
}

func TestBillRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BillRepositoryTestSuite))
}

// Test Cases
func (suite *BillRepositoryTestSuite) TestCreateBill_Success() {
	// Preparation
	dummyBill := model.Bill{
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
					Id:   "1",
					Name: "Cuci + Setrika",
				},
				Qty:   1,
				Price: 20000,
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// EKSPEKTASI
	suite.sqlmock.ExpectBegin()

	rows := sqlmock.NewRows([]string{"id", "bill_date", "created_at", "updated_at"}).AddRow(dummyBill.Id, dummyBill.BillDate, dummyBill.CreatedAt, dummyBill.UpdatedAt)

	suite.sqlmock.ExpectQuery("INSERT INTO bills").WillReturnRows(rows)

	for _, v := range dummyBill.BillDetails {
		rows := sqlmock.NewRows([]string{"id", "qty", "price", "created_at", "updated_at"}).AddRow(v.Id, v.Qty, v.Price, v.CreatedAt, v.UpdatedAt)
		suite.sqlmock.ExpectQuery("INSERT INTO bill_details").WillReturnRows(rows)
	}

	suite.sqlmock.ExpectCommit()

	// EKSEKUSI / AKTUAL
	actual, err := suite.repo.Create(dummyBill)

	// ASSERTION
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyBill.Id, actual.Id)
}

func (suite *BillRepositoryTestSuite) TestCreateBill_Fail() {
	dummyBill := model.Bill{
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
					Id:   "1",
					Name: "Cuci + Setrika",
				},
				Qty:   1,
				Price: 20000,
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// BEGIN
	suite.sqlmock.ExpectBegin().WillReturnError(errors.New("error begin"))
	_, err := suite.repo.Create(dummyBill)
	assert.Error(suite.T(), err)

	// INSERT BILLS
	suite.sqlmock.ExpectBegin() // reset begin
	suite.sqlmock.ExpectQuery("INSERT INTO bills").WillReturnError(errors.New("insert failed"))
	_, err = suite.repo.Create(dummyBill)
	assert.Error(suite.T(), err)

	// SUCCESS
	suite.sqlmock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "bill_date", "created_at", "updated_at"}).AddRow(dummyBill.Id, dummyBill.BillDate, dummyBill.CreatedAt, dummyBill.UpdatedAt)
	suite.sqlmock.ExpectQuery("INSERT INTO bills").WillReturnRows(rows)

	// INSERT DETAILS
	for _, v := range dummyBill.BillDetails {
		fmt.Println(v)
		suite.sqlmock.ExpectBegin() // reset begin
		suite.sqlmock.ExpectQuery("INSERT INTO bill_details").WillReturnError(errors.New("insert failed"))
		_, err = suite.repo.Create(dummyBill)
		assert.Error(suite.T(), err)
	}
}
