package repository_datamart

import (
	"context"
	"database/sql"
	"moonlay/model/domain"
)

type RepositoryDatamart interface {
	GetDatamar1(ctx context.Context, tx *sql.Tx, interval int) []domain.Datamart1
	GetDatamar2(ctx context.Context, tx *sql.Tx, interval int) []domain.Datamart2
	GetDatamar3(ctx context.Context, tx *sql.Tx, interval int) []domain.Datamart3
	GetTotalData(ctx context.Context, tx *sql.Tx) int
}
