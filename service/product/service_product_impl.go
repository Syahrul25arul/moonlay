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
	/*
	* setup tx transaction from db
	* commit or rollback at the end of this method
	 */
	tx, err := s.DB.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	// * set struct products for send to repository
	product := &domain.Products{}
	repo := s.repositoryProduct

	/*
	* iterate data array products for excel file
	* and send to repository for save to db
	 */
	helper.ExtractMultiDimensitinString(data, func(s []string) {
		product.ConvertDataStringToProducts(s)
		repo.Save(ctx, tx, *product)
	})
}
