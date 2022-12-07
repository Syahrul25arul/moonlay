package domain

type Customers struct {
	Id           int
	CustomerId   string `db:"customer_id"`
	CustomerName string `db:"customer_name"`
	Status       string `db:"status"`
}
