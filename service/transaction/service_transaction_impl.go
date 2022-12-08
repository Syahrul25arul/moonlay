package service_transaction

import (
	"context"
	"database/sql"
	"moonlay/helper"
	"moonlay/model/domain"
	repositor_transaction "moonlay/repository/transactions"
)

type serviceTransactionImpl struct {
	repository repositor_transaction.RepositoryTransaction
	DB         *sql.DB
}

func NewServiceTransaction(repository repositor_transaction.RepositoryTransaction, db *sql.DB) ServiceTransaction {
	return &serviceTransactionImpl{
		repository: repository,
		DB:         db,
	}
}

func (s *serviceTransactionImpl) CreateTransactionFromFile(ctx context.Context, data [][]string) {
	/*
	* setup tx transaction from db
	* commit or rollback at the end of this method
	 */
	tx, err := s.DB.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	// * set struct transaction for send to repository
	transaction := &domain.Transactions{}
	repo := s.repository

	/*
	* iterate data array transactions for excel file
	* and send to repository for save to db
	 */
	helper.ExtractMultiDimensitinString(data, func(s []string) {
		transaction.ConvertDataStringToTransactions(s)
		repo.Save(context.Background(), tx, *transaction)
	})
}
