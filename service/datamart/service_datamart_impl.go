package service_datamart

import (
	"context"
	"database/sql"
	"moonlay/helper"
	"moonlay/model/domain"
	repository_datamart "moonlay/repository/datamart"
)

type serviceDatamartImpl struct {
	repository repository_datamart.RepositoryDatamart
	db         *sql.DB
}

func NewServiceDatamart(repository repository_datamart.RepositoryDatamart, db *sql.DB) ServiceDatamart {
	return &serviceDatamartImpl{
		repository: repository,
		db:         db,
	}
}

func (s *serviceDatamartImpl) GetDatamar1(ctx context.Context, interval int) []domain.Datamart1 {
	tx, err := s.db.Begin()
	helper.PanicIFError(err)

	defer helper.CommitOrRollback(tx)
	return s.repository.GetDatamar1(ctx, tx, interval)
}

func (s *serviceDatamartImpl) GetTotalData(ctx context.Context) int {
	tx, err := s.db.Begin()
	helper.PanicIFError(err)

	defer helper.CommitOrRollback(tx)
	return s.repository.GetTotalData(ctx, tx)
}
func (s *serviceDatamartImpl) GetDatamar2(ctx context.Context, interval int) []domain.Datamart2 {
	tx, err := s.db.Begin()
	helper.PanicIFError(err)

	defer helper.CommitOrRollback(tx)
	return s.repository.GetDatamar2(ctx, tx, interval)
}

func (s *serviceDatamartImpl) GetDatamar3(ctx context.Context, interval int) []domain.Datamart3 {
	tx, err := s.db.Begin()
	helper.PanicIFError(err)

	defer helper.CommitOrRollback(tx)
	return s.repository.GetDatamar3(ctx, tx, interval)
}
