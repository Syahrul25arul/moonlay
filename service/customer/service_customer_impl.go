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
	/*
	* setup tx transaction from db
	* commit or rollback at the end of this method
	 */
	tx, err := s.DB.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	// * set struct customer for send to repository
	var customer = &domain.Customers{}
	repo := s.repositoryCustomer

	/*
	* iterate data array customers for excel file
	* and send to repository for save to db
	 */
	helper.ExtractMultiDimensitinString(data, func(s []string) {
		customer.ConvertDataStringToCustomers(s)
		repo.Save(context.Background(), tx, *customer)
	})
}
