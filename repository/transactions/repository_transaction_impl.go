package repository_transaction

import (
	"context"
	"database/sql"
	"moonlay/helper"
	"moonlay/model/domain"
)

type repositoryTransactionImpl struct {
}

func NewRepositoryTransaction() RepositoryTransaction {
	return &repositoryTransactionImpl{}
}

func (r *repositoryTransactionImpl) Save(ctx context.Context, tx *sql.Tx, transaction domain.Transactions) domain.Transactions {
	SQL := "insert into transactions(transaction_id,transaction_uuid,rel_uuid,buyer_id,seller_id,product_id,price,volume,value,transaction_date,entry_date,confirm_date,complete_data_buyer,complete_data_seller,buy_sell,is_amend,is_cancel,confirm_status,complete_status_buyer,complete_status_seller,status) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21) returning id"

	LastInserId := 0

	err := tx.QueryRowContext(ctx, SQL, transaction.TransactionId, transaction.TransactionUuid, transaction.RelUuid, transaction.BuyerId, transaction.SellerId, transaction.ProductId, transaction.Price, transaction.Volume, transaction.Value, transaction.TransactionDate, transaction.EntryDate, transaction.ConfirmDate, transaction.CompleteDataBuyer, transaction.CompleteDataSeller, transaction.BuySell, transaction.IsAmmend, transaction.IsCancel, transaction.ConfirmStatus, transaction.CompleteStatusBuyer, transaction.CompleteStatusSeller, transaction.Status).Scan(&LastInserId)
	helper.PanicIFError(err)

	transaction.Id = int(LastInserId)
	return transaction
}
