package service_customer

import "context"

type ServiceCustomer interface {
	CreateCustomerFromFile(ctx context.Context, data [][]string)
}
