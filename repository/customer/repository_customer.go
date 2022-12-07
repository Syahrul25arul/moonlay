package repository_customer

import (
	"context"
	"database/sql"
	"moonlay/model/domain"
)

type RepositoryCustomer interface {
	Save(ctx context.Context, tx *sql.Tx, customers domain.Customers) domain.Customers
}
