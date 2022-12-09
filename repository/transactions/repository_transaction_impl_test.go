package repository_transaction

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

func TestNewRepositoryTransaction(t *testing.T) {
	repository := NewRepositoryTransaction()
	reflection := reflect.TypeOf(repository)
	assert.Equal(t, "*repository_transaction.repositoryTransactionImpl", reflection.String())
}

func Test_repositoryTransactionImpl_Save(t *testing.T) {
	db := database.NewDB()
	tx, err := db.Begin()
	helper.PanicIFError(err)

	// this function for handle panic helper.CommitOrRollback
	defer func() {
		recover()
	}()
	defer helper.CommitOrRollback(tx)

	database.TruncateAllTable(db)

	repository := NewRepositoryTransaction()

	tests := []struct {
		name     string
		r        RepositoryTransaction
		want     domain.Transactions
		expected domain.Transactions
		err      string
	}{
		{
			name: "save transactions success",
			r:    repository,
			want: domain.Transactions{
				TransactionId:        "TR00-1",
				TransactionUuid:      "10fc653e-9f40-449c-9228-a070c0724ae1",
				RelUuid:              "026a76a2-7bab-4244-97df-af9b7dd150a4",
				BuyerId:              "CTM00-89",
				SellerId:             "CTM00-130",
				ProductId:            "PRD00-894",
				Price:                33,
				Volume:               94,
				Value:                310200000,
				TransactionDate:      helper.ParseTime("2/23/22 22:08"),
				EntryDate:            helper.ParseTime("2/24/22 12:03"),
				ConfirmDate:          helper.ParseTime("2/25/22 12:10"),
				CompleteDataBuyer:    helper.ParseTime("2/26/22 15:50"),
				CompleteDataSeller:   helper.ParseTime("2/27/22 12:23"),
				BuySell:              "B",
				IsAmmend:             "N",
				IsCancel:             "N",
				ConfirmStatus:        "CONFIRMED",
				CompleteStatusBuyer:  "COMPLETE1",
				CompleteStatusSeller: "COMPLETE1",
				Status:               "M",
			},
			expected: domain.Transactions{
				Id:                   1,
				TransactionId:        "TR00-1",
				TransactionUuid:      "10fc653e-9f40-449c-9228-a070c0724ae1",
				RelUuid:              "026a76a2-7bab-4244-97df-af9b7dd150a4",
				BuyerId:              "CTM00-89",
				SellerId:             "CTM00-130",
				ProductId:            "PRD00-894",
				Price:                33,
				Volume:               94,
				Value:                310200000,
				TransactionDate:      helper.ParseTime("2/23/22 22:08"),
				EntryDate:            helper.ParseTime("2/24/22 12:03"),
				ConfirmDate:          helper.ParseTime("2/25/22 12:10"),
				CompleteDataBuyer:    helper.ParseTime("2/26/22 15:50"),
				CompleteDataSeller:   helper.ParseTime("2/27/22 12:23"),
				BuySell:              "B",
				IsAmmend:             "N",
				IsCancel:             "N",
				ConfirmStatus:        "CONFIRMED",
				CompleteStatusBuyer:  "COMPLETE1",
				CompleteStatusSeller: "COMPLETE1",
				Status:               "M",
			},
			err: "",
		},
		{
			name:     "save transactions failed",
			r:        repository,
			want:     domain.Transactions{},
			expected: domain.Transactions{},
			err:      "invalid input value for enum status: \"\"",
		},
		{
			name:     "save transactions failed lose db connection",
			r:        repository,
			want:     domain.Transactions{},
			expected: domain.Transactions{},
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
