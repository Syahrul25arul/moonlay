package repository_product

import (
	"context"
	"database/sql"
	"moonlay/model/domain"
)

type RepositoryProduct interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products
}
