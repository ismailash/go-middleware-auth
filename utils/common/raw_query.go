package common

const (
	CreateBill        = `INSERT INTO bills (bill_date,customer_id,user_id,updated_at) VALUES ($1,$2,$3,$4) RETURNING id,bill_date,created_at,updated_at`
	CreateBillDetail  = `INSERT INTO bill_details (bill_id,product_id,qty,price,updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING id,qty,price,created_at,updated_at`
	GetBillById       = `SELECT b.id, b.bill_date, c.id, c.name, c.phone_number, c.address, c.created_at, c.updated_at, u.id, u.name, u.email, u.username, u.role, u.created_at, u.updated_at, b.created_at, b.updated_at FROM bills b JOIN customers c ON c.id = b.customer_id JOIN users u ON u.id = b.user_id WHERE b.id = $1`
	GetBillDetailById = ` SELECT bd.id, p.id, p."name", p.price, p."type", p.created_at, p.updated_at, bd.qty, bd.price, bd.created_at, bd.updated_at FROM bill_details bd JOIN bills b ON 	b.id = bd.bill_id JOIN products p ON p.id = bd.product_id WHERE b.id = $1 `
)
