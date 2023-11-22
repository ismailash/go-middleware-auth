package dto

import "enigmacamp.com/be-enigma-laundry/model"

type BillRequestDto struct {
	Id          string             `json:"id"`
	CustomerId  string             `json:"customerId"`
	UserId      string             `json:"userId"`
	BillDetails []model.BillDetail `json:"billDetails"`
}

// Example Payload Request
/*

	{
		"customerId": "",
		"billDetails": [
			{
				"product": { "id": "" },
				"qty": 1
			},
			{
				"product": { "id": "" },
				"qty": 1
			}
		]
	}
*/
