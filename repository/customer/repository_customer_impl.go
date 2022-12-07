package repository_customer

import (
	"context"
	"database/sql"
	"moonlay/helper"
	"moonlay/model/domain"
)

type RepositoryCustomerImpl struct {
}

func NewRepositoryCustomer() RepositoryCustomer {
	return &RepositoryCustomerImpl{}
}

func (r *RepositoryCustomerImpl) Save(ctx context.Context, tx *sql.Tx, customers domain.Customers) domain.Customers {
	SQL := "insert into customers(customer_id,customer_name,status) values ($1,$2,$3) returning id"
	LastInserId := 0
	err := tx.QueryRowContext(ctx, SQL, customers.CustomerId, customers.CustomerName, customers.Status).Scan(&LastInserId)
	helper.PanicIFError(err)

	customers.Id = int(LastInserId)
	return customers
}
