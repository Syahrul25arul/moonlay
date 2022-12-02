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
	SQL := "insert into customers(customer_id,customer_name,status) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, customers.CustomerId, customers.CustomerName, customers.Status)
	helper.PanicIFError(err)

	id, err := result.LastInsertId()
	helper.PanicIFError(err)

	customers.Id = int(id)
	return customers
}
