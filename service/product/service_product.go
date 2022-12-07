package service_product

import "context"

type ServiceProduct interface {
	CreateProductFromFile(ctx context.Context, data [][]string)
}
