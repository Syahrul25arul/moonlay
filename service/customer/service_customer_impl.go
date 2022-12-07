package service_customer

import (
	"context"
	"database/sql"
	"moonlay/helper"
	"moonlay/model/domain"
	repository_customer "moonlay/repository/customer"
)

type ServiceCustomerImpl struct {
	repositoryCustomer repository_customer.RepositoryCustomer
	DB                 *sql.DB
}

func NewServiceCustomer(repositoryCustomer repository_customer.RepositoryCustomer, tx *sql.DB) ServiceCustomer {
	return &ServiceCustomerImpl{repositoryCustomer: repositoryCustomer, DB: tx}
}

func (s *ServiceCustomerImpl) CreateCustomerFromFile(ctx context.Context, data [][]string) {
	tx, err := s.DB.Begin()

	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	var customer = domain.Customers{}

	for j, valueSlice := range data {
		if j == 0 {
			continue
		}
		for i, valueString := range valueSlice {

			switch {
			case (i+1)%3 == 0:
				customer.CustomerName = valueString
			case (i+1)%4 == 0:
				customer.Status = valueString
				s.repositoryCustomer.Save(ctx, tx, customer)
			default:
				customer.CustomerId = valueString
			}

		}
	}
}
