package repository

import (
	"database/sql"
	"enigmacamp.com/be-enigma-laundry/utils/common"
	"time"

	"enigmacamp.com/be-enigma-laundry/model"
)

type BillRepository interface {
	Create(payload model.Bill) (model.Bill, error)
	Get(id string) (model.Bill, error)
}

type billRepository struct {
	db *sql.DB
}

func (b *billRepository) Create(payload model.Bill) (model.Bill, error) {
	tx, err := b.db.Begin()
	if err != nil {
		return model.Bill{}, err
	}

	var bill model.Bill
	err = tx.QueryRow(common.CreateBill,
		time.Now(), payload.Customer.Id, payload.User.Id, time.Now()).
		Scan(
			&bill.Id,
			&bill.BillDate,
			&bill.CreatedAt,
			&bill.UpdatedAt,
		)

	if err != nil {
		return model.Bill{}, tx.Rollback()
	}

	var billDetails []model.BillDetail
	for _, v := range payload.BillDetails {
		var billDetail model.BillDetail
		err = tx.QueryRow(common.CreateBillDetail,
			bill.Id, v.Product.Id, v.Qty, v.Price, time.Now()).
			Scan(
				&billDetail.Id,
				&billDetail.Qty,
				&billDetail.Price,
				&billDetail.CreatedAt,
				&billDetail.UpdatedAt,
			)

		if err != nil {
			return model.Bill{}, tx.Rollback()
		}
		billDetail.Product = v.Product
		billDetails = append(billDetails, billDetail)
	}
	bill.Customer = payload.Customer
	bill.User = payload.User
	bill.BillDetails = billDetails
	if err := tx.Commit(); err != nil {
		return model.Bill{}, err
	}

	return bill, nil
}

func (b *billRepository) Get(id string) (model.Bill, error) {
	var bill model.Bill
	err := b.db.QueryRow(common.GetBillById, id).Scan(
		&bill.Id,
		&bill.BillDate,
		&bill.Customer.Id,
		&bill.Customer.Name,
		&bill.Customer.PhoneNumber,
		&bill.Customer.Address,
		&bill.Customer.CreatedAt,
		&bill.Customer.UpdatedAt,
		&bill.User.Id,
		&bill.User.Name,
		&bill.User.Email,
		&bill.User.Username,
		&bill.User.Role,
		&bill.User.CreatedAt,
		&bill.User.UpdatedAt,
		&bill.CreatedAt,
		&bill.UpdatedAt,
	)

	if err != nil {
		return model.Bill{}, err
	}

	var billDetails []model.BillDetail
	rows, err := b.db.Query(common.GetBillDetailById, bill.Id)

	if err != nil {
		return model.Bill{}, err
	}

	for rows.Next() {
		var billDetail model.BillDetail
		rows.Scan(
			&billDetail.Id,
			&billDetail.Product.Id,
			&billDetail.Product.Name,
			&billDetail.Product.Price,
			&billDetail.Product.Type,
			&billDetail.Product.CreatedAt,
			&billDetail.Product.UpdatedAt,
			&billDetail.Qty,
			&billDetail.Price,
			&billDetail.CreatedAt,
			&billDetail.UpdatedAt,
		)

		billDetails = append(billDetails, billDetail)
	}
	bill.BillDetails = billDetails
	return bill, nil
}

func NewBillRepository(db *sql.DB) BillRepository {
	return &billRepository{db: db}
}
