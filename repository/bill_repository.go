package repository

import (
	"database/sql"
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
	err = tx.QueryRow(`INSERT INTO bills (bill_date,customer_id,user_id,updated_at) VALUES ($1,$2,$3,$4) RETURNING id,bill_date,created_at,updated_at`,
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
		err = tx.QueryRow(`INSERT INTO bill_details (bill_id,product_id,qty,price,updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING id,qty,price,created_at,updated_at`,
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
	err := b.db.QueryRow(`
	SELECT
		b.id,
		b.bill_date,
		c.id,
		c.name,
		c.phone_number,
		c.address,
		c.created_at,
		c.updated_at,
		u.id,
		u.name,
		u.email,
		u.username,
		u.role,
		u.created_at,
		u.updated_at,
		b.created_at,
		b.updated_at
	FROM
		bills b
	JOIN customers c ON
		c.id = b.customer_id
	JOIN users u ON
		u.id = b.user_id
	WHERE b.id = $1
	`, id).Scan(
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
	rows, err := b.db.Query(`
	SELECT
		bd.id,
		p.id,
		p."name",
		p.price,
		p."type",
		p.created_at,
		p.updated_at,
		bd.qty,
		bd.price,
		bd.created_at,
		bd.updated_at
	FROM
		bill_details bd
	JOIN bills b ON
		b.id = bd.bill_id
	JOIN products p ON
		p.id = bd.product_id
	WHERE
		b.id = $1
	`, bill.Id)

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
