package service_transaction

import (
	"context"
	"database/sql"
	"moonlay/helper"
	"moonlay/model/domain"
	repositor_transaction "moonlay/repository/transactions"
	"strconv"
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
	tx, err := s.DB.Begin()
	helper.PanicIFError(err)

	defer helper.CommitOrRollback(tx)

	transaction := domain.Transactions{}

	for i, valueSlice := range data {
		if i == 0 {
			continue
		}

		index := 1
		for _, cellValue := range valueSlice {
			switch {
			case index == 3:
				transaction.TransactionUuid = cellValue
			case index == 4:
				transaction.RelUuid = cellValue
			case index == 5:
				transaction.BuyerId = cellValue
			case index == 6:
				transaction.SellerId = cellValue
			case index == 7:
				transaction.ProductId = cellValue
			case index == 8:
				price, err := strconv.Atoi(cellValue)
				helper.PanicIFError(err)

				transaction.Price = price
			case index == 9:
				volume, err := strconv.Atoi(cellValue)
				helper.PanicIFError(err)
				transaction.Volume = volume
			case index == 10:
				value, err := strconv.ParseUint(cellValue, 10, 64)
				helper.PanicIFError(err)
				transaction.Value = value
			case index == 11:
				transaction.TransactionDate = helper.ParseTime(cellValue)
			case index == 12:
				transaction.EntryDate = helper.ParseTime(cellValue)
			case index == 13:
				transaction.ConfirmDate = helper.ParseTime(cellValue)
			case index == 14:
				transaction.CompleteDataBuyer = helper.ParseTime(cellValue)
			case index == 15:
				transaction.CompleteDataSeller = helper.ParseTime(cellValue)
			case index == 16:
				transaction.BuySell = cellValue
			case index == 17:
				transaction.IsAmmend = cellValue
			case index == 18:
				transaction.IsCancel = cellValue
			case index == 19:
				transaction.ConfirmStatus = cellValue
			case index == 20:
				transaction.CompleteStatusBuyer = cellValue
			case index == 21:
				transaction.CompleteStatusSeller = cellValue
			case index == 22:
				transaction.Status = cellValue
				s.repository.Save(context.Background(), tx, transaction)
				index = 1
			default:
				transaction.TransactionId = cellValue
			}
			index++

		}
	}
}
