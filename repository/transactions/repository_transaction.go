package repository_transaction

import (
	"context"
	"database/sql"
	"moonlay/model/domain"
)

type RepositoryTransaction interface {
	Save(ctx context.Context, tx *sql.Tx, transaction domain.Transactions) domain.Transactions
}
