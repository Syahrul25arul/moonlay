package repository_product

import (
	"context"
	"database/sql"
	"moonlay/helper"
	"moonlay/model/domain"
)

type RepositoryProductImpl struct {
}

func NewRepositoryProduct() RepositoryProduct {
	return &RepositoryProductImpl{}
}

func (r *RepositoryProductImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products {
	SQL := "insert into products(product_id,product_name,currency_code,status) values ($1,$2,$3,$4) returning id"
	LastInserId := 0
	err := tx.QueryRowContext(ctx, SQL, product.ProductId, product.ProductName, product.CurrencyCode, product.Status).Scan(&LastInserId)
	helper.PanicIFError(err)

	product.Id = int(LastInserId)
	return product
}
