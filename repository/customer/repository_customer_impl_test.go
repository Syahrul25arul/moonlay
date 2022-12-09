package repository_customer

import (
	"context"
	"fmt"
	"moonlay/database"
	"moonlay/helper"
	"moonlay/model/domain"
	"reflect"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestNewRepositoryCustomer(t *testing.T) {
	repository := NewRepositoryCustomer()

	reflection := reflect.TypeOf(repository)

	assert.Equal(t, "*repository_customer.RepositoryCustomerImpl", reflection.String())
}

func TestRepositoryCustomerImpl_Save(t *testing.T) {
	db := database.NewDB()
	tx, err := db.Begin()
	helper.PanicIFError(err)

	// this function for handle panic helper.CommitOrRollback
	defer func() {
		recover()
	}()
	defer helper.CommitOrRollback(tx)

	database.TruncateAllTable(tx)

	repository := NewRepositoryCustomer()

	tests := []struct {
		name     string
		r        RepositoryCustomer
		want     domain.Customers
		expected domain.Customers
		err      string
	}{
		{
			name: "save customer success",
			r:    repository,
			want: domain.Customers{
				CustomerId:   "CTM-001",
				CustomerName: "hendrik",
				Status:       "Y",
			},
			expected: domain.Customers{
				Id:           1,
				CustomerId:   "CTM-001",
				CustomerName: "hendrik",
				Status:       "Y",
			},
			err: "",
		},
		{
			name:     "save customer failed",
			r:        repository,
			want:     domain.Customers{},
			expected: domain.Customers{},
			err:      "invalid input value for enum status: \"\"",
		},
		{
			name:     "save customer failed connection close",
			r:        repository,
			want:     domain.Customers{},
			expected: domain.Customers{},
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
			fmt.Println(tt.name)

			assert.Equal(t, tt.expected, result)

		})
	}
}
