package service_product

import (
	"context"
	"database/sql"
	"moonlay/helper"
	"moonlay/model/domain"
	repository_product "moonlay/repository/product"
)

type serviceProductImpl struct {
	repositoryProduct repository_product.RepositoryProduct
	DB                *sql.DB
}

func NewServiceProduct(repository repository_product.RepositoryProduct, db *sql.DB) ServiceProduct {
	return &serviceProductImpl{repositoryProduct: repository, DB: db}
}

func (s *serviceProductImpl) CreateProductFromFile(ctx context.Context, data [][]string) {
	tx, err := s.DB.Begin()

	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Products{}

	for i, valueSlice := range data {
		if i == 0 {
			continue
		}

		for j, valueCell := range valueSlice {

			switch {
			case (j+1)%3 == 0:
				product.ProductName = valueCell
			case (j+1)%4 == 0:
				product.CurrencyCode = valueCell
			case (j+1)%5 == 0:
				product.Status = valueCell
				s.repositoryProduct.Save(context.Background(), tx, product)
			default:
				product.ProductId = valueCell
			}
		}
	}
}
