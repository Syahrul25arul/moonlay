package repository_product

import (
	"context"
	"moonlay/database"
	"moonlay/helper"
	"moonlay/model/domain"
	"reflect"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestNewRepositoryProduct(t *testing.T) {
	repository := NewRepositoryProduct()

	reflection := reflect.TypeOf(repository)

	assert.Equal(t, "*repository_product.RepositoryProductImpl", reflection.String())
}

func TestRepositoryProductImpl_Save(t *testing.T) {
	db := database.NewDB()
	tx, err := db.Begin()
	helper.PanicIFError(err)

	// this function for handle panic helper.CommitOrRollback
	defer func() {
		recover()
	}()
	defer helper.CommitOrRollback(tx)

	database.TruncateAllTable(db)

	repository := NewRepositoryProduct()

	tests := []struct {
		name     string
		r        RepositoryProduct
		want     domain.Products
		expected domain.Products
		err      string
	}{
		{
			name: "save customers success",
			r:    repository,
			want: domain.Products{
				ProductId:    "CTM-001",
				ProductName:  "keyboard logitec",
				CurrencyCode: "LDR",
				Status:       "Y",
			},
			expected: domain.Products{
				Id:           1,
				ProductId:    "CTM-001",
				ProductName:  "keyboard logitec",
				CurrencyCode: "LDR",
				Status:       "Y",
			},
			err: "",
		},
		{
			name:     "save customers failed",
			r:        repository,
			want:     domain.Products{},
			expected: domain.Products{},
			err:      "invalid input value for enum status: \"\"",
		},
		{
			name:     "save customers failed close db",
			r:        repository,
			want:     domain.Products{},
			expected: domain.Products{},
			err:      "current transaction is aborted, commands ignored until end of transaction block",
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					errs := r.(*pq.Error)
					assert.Equal(t, tt.err, errs.Message)
				}
			}()

			if i == 2 {
				db.Close()
			}

			result := tt.r.Save(context.Background(), tx, tt.want)

			assert.Equal(t, tt.expected, result)
		})
	}
}
