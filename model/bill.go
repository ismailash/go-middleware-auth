package model

import "time"

type Bill struct {
	Id          string       `json:"id"`
	BillDate    time.Time    `json:"billDate"`
	Customer    Customer     `json:"customer"`
	User        User         `json:"user"` // employee
	BillDetails []BillDetail `json:"billDetails"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type BillDetail struct {
	Id        string    `json:"id"`
	BillId    string    `json:"billId,omitempty"`
	Product   Product   `json:"product"`
	Qty       int       `json:"qty"`
	Price     int       `json:"price"` // automatically get from product
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
