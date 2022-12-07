package repository_datamart

import (
	"context"
	"database/sql"
	"moonlay/helper"
	"moonlay/model/domain"
)

type repositoryDatamartImpl struct {
}

func NewRepositoryDatamart() RepositoryDatamart {
	return &repositoryDatamartImpl{}
}

func (r *repositoryDatamartImpl) GetDatamar1(ctx context.Context, tx *sql.Tx, interval int) []domain.Datamart1 {
	var offset int
	if interval == 0 {
		offset = 0
	} else {
		offset = (interval * 100)
	}
	sql := "select * from datamart limit 100 offset $1"
	rows, err := tx.QueryContext(ctx, sql, offset)
	helper.PanicIFError(err)

	defer rows.Close()

	var datamarts []domain.Datamart1
	for rows.Next() {

		datamart := domain.Datamart1{}
		err := rows.Scan(
			&datamart.TransactionId,
			&datamart.BuyerId,
			&datamart.BuyerName,
			&datamart.SellerId,
			&datamart.SellerName,
			&datamart.ProductId,
			&datamart.ProductName,
			&datamart.Currency,
			&datamart.Price,
			&datamart.Volume,
			&datamart.Value,
			&datamart.TransactionDate,
			&datamart.TransactionMonth,
			&datamart.TransactionYear,
			&datamart.EntryDate,
			&datamart.EntryMonth,
			&datamart.EntryYear,
			&datamart.Buysell,
			&datamart.ConfirmStatus,
			&datamart.CompleteStatusBuyer,
			&datamart.CompleteStatusSeller,
		)

		helper.PanicIFError(err)
		datamarts = append(datamarts, datamart)
	}
	return datamarts
}

func (r *repositoryDatamartImpl) GetTotalData(ctx context.Context, tx *sql.Tx) int {
	sql := "select count(d.transaction_id) from datamart d"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIFError(err)

	defer rows.Close()

	var total int
	for rows.Next() {
		err := rows.Scan(&total)
		helper.PanicIFError(err)
	}
	return total
}
func (r *repositoryDatamartImpl) GetDatamar2(ctx context.Context, tx *sql.Tx, interval int) []domain.Datamart2 {
	var offset int
	if interval == 0 {
		offset = 0
	} else {
		offset = (interval * 100)
	}
	sql := "select * from datamart2 limit 100 offset $1"
	rows, err := tx.QueryContext(ctx, sql, offset)
	helper.PanicIFError(err)

	defer rows.Close()

	var datamarts []domain.Datamart2
	for rows.Next() {
		datamart := domain.Datamart2{}
		err := rows.Scan(
			&datamart.ProductId,
			&datamart.ProductName,
			&datamart.Price,
			&datamart.Volume,
			&datamart.Value,
			&datamart.TransactionDate,
			&datamart.EntryDate,
		)

		helper.PanicIFError(err)
		datamarts = append(datamarts, datamart)
	}
	return datamarts
}

func (r *repositoryDatamartImpl) GetDatamar3(ctx context.Context, tx *sql.Tx, interval int) []domain.Datamart3 {
	var offset int
	if interval == 0 {
		offset = 0
	} else {
		offset = (interval * 100)
	}
	sql := "select * from datamart3 limit 100 offset $1"
	rows, err := tx.QueryContext(ctx, sql, offset)
	helper.PanicIFError(err)

	defer rows.Close()

	var datamarts []domain.Datamart3
	for rows.Next() {
		datamart := domain.Datamart3{}

		err := rows.Scan(
			&datamart.CustomerId,
			&datamart.CustomerName,
			&datamart.Price,
			&datamart.Volume,
			&datamart.Value,
			&datamart.TransactionDate,
			&datamart.EntryDate,
		)

		helper.PanicIFError(err)
		datamarts = append(datamarts, datamart)
	}
	return datamarts
}
