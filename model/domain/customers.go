package domain

type Customers struct {
	Id           int
	CustomerId   string `db:"customer_id"`
	CustomerName string `db:"customer_name"`
	Status       string `db:"status"`
}

func (c *Customers) ConvertDataStringToCustomers(data []string) {
	c.CustomerId = data[1]
	c.CustomerName = data[2]
	c.Status = data[3]

}
