package domain

type Products struct {
	Id           int
	ProductId    string
	ProductName  string
	CurrencyCode string
	Status       string
}

func (p *Products) ConvertDataStringToProducts(data []string) {
	p.ProductId = data[1]
	p.ProductName = data[3]
	p.CurrencyCode = data[3]
	p.Status = data[4]
}
