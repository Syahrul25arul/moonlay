package service_transaction

import "context"

type ServiceTransaction interface {
	CreateTransactionFromFile(ctx context.Context, data [][]string)
}
