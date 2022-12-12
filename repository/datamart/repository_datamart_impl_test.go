package repository_datamart

import (
	"context"
	"moonlay/database"
	"moonlay/helper"
	"moonlay/model/domain"
	repository_customer "moonlay/repository/customer"
	repository_product "moonlay/repository/product"
	repository_transaction "moonlay/repository/transactions"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDataDummy() {
	db := database.NewDB()
	defer db.Close()
	tx, err := db.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	repoProduct := repository_product.NewRepositoryProduct()
	repoCustomer := repository_customer.NewRepositoryCustomer()
	repoTransaction := repository_transaction.NewRepositoryTransaction()

	product := domain.Products{
		ProductId:    "PRD-001",
		ProductName:  "keyboard logitec",
		CurrencyCode: "LDR",
		Status:       "Y",
	}
	customer := domain.Customers{
		CustomerId:   "CTM-001",
		CustomerName: "hendrik",
		Status:       "Y",
	}
	customer2 := domain.Customers{
		CustomerId:   "CTM-002",
		CustomerName: "array",
		Status:       "Y",
	}
	transaction := domain.Transactions{
		TransactionId:        "TR00-1",
		TransactionUuid:      "10fc653e-9f40-449c-9228-a070c0724ae1",
		RelUuid:              "026a76a2-7bab-4244-97df-af9b7dd150a4",
		BuyerId:              "CTM-001",
		SellerId:             "CTM-002",
		ProductId:            "PRD-001",
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
	}

	repoProduct.Save(context.Background(), tx, product)
	repoCustomer.Save(context.Background(), tx, customer)
	repoCustomer.Save(context.Background(), tx, customer2)
	repoTransaction.Save(context.Background(), tx, transaction)

}

func TestNewRepositoryDatamart(t *testing.T) {
	repository := NewRepositoryDatamart()

	reflection := reflect.TypeOf(repository)

	assert.Equal(t, "*repository_datamart.repositoryDatamartImpl", reflection.String())
}

func Test_repositoryDatamartImpl_GetDatamar1(t *testing.T) {
	db := database.NewDB()
	tx, err := db.Begin()
	helper.PanicIFError(err)

	// this function for handle panic helper.CommitOrRollback
	defer func() {
		recover()
	}()
	defer helper.CommitOrRollback(tx)
	database.TruncateAllTable(db)

	repository := NewRepositoryDatamart()

	tests := []struct {
		name     string
		r        RepositoryDatamart
		want     int
		expected []domain.Datamart1
		err      string
	}{
		{
			name:     "get datamart null data",
			r:        repository,
			want:     0,
			expected: []domain.Datamart1(nil),
		},
		{
			name: "get datamart success",
			r:    repository,
			want: 0,
			expected: []domain.Datamart1{
				{
					TransactionId:        "TR00-1",
					BuyerId:              "CTM-001",
					BuyerName:            "hendrik",
					SellerId:             "CTM-002",
					SellerName:           "array",
					ProductId:            "PRD-001",
					ProductName:          "keyboard logitec",
					Currency:             "LDR",
					Price:                33,
					Volume:               94,
					Value:                310200000,
					TransactionDate:      helper.ParseTime("2/23/22 22:08"),
					TransactionMonth:     2,
					TransactionYear:      2022,
					EntryDate:            helper.ParseTime("2/24/22 12:03"),
					EntryMonth:           2,
					EntryYear:            2022,
					Buysell:              "B",
					ConfirmStatus:        "CONFIRMED",
					CompleteStatusBuyer:  "COMPLETE1",
					CompleteStatusSeller: "COMPLETE1",
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 0 {
				result := tt.r.GetDatamar1(context.Background(), tx, 0)
				assert.Equal(t, tt.expected, result)
				return
			}
			if i == 1 {
				createDataDummy()
			}

			result := tt.r.GetDatamar1(context.Background(), tx, 0)

			datamart := result[0]
			datamart.TransactionDate = helper.ParseTime(helper.ParseTimeToString(datamart.TransactionDate))
			datamart.EntryDate = helper.ParseTime(helper.ParseTimeToString(datamart.EntryDate))
			result[0] = datamart
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_repositoryDatamartImpl_GetTotalData(t *testing.T) {
	db := database.NewDB()
	tx, err := db.Begin()
	helper.PanicIFError(err)

	// this function for handle panic helper.CommitOrRollback
	defer func() {
		recover()
	}()
	defer helper.CommitOrRollback(tx)
	database.TruncateAllTable(db)
	createDataDummy()

	repository := NewRepositoryDatamart()

	tests := []struct {
		name     string
		r        RepositoryDatamart
		expected int
		err      string
	}{
		{
			name:     "get datamart null data",
			r:        repository,
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.r.GetTotalData(context.Background(), tx)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_repositoryDatamartImpl_GetDatamar2(t *testing.T) {
	db := database.NewDB()
	tx, err := db.Begin()
	helper.PanicIFError(err)

	// this function for handle panic helper.CommitOrRollback
	defer func() {
		recover()
	}()
	defer helper.CommitOrRollback(tx)
	database.TruncateAllTable(db)

	repository := NewRepositoryDatamart()

	tests := []struct {
		name     string
		r        RepositoryDatamart
		want     int
		expected []domain.Datamart2
		err      string
	}{
		{
			name:     "get datamart null data",
			r:        repository,
			want:     0,
			expected: []domain.Datamart2(nil),
		},
		{
			name: "get datamart success",
			r:    repository,
			want: 0,
			expected: []domain.Datamart2{
				{
					ProductId:       "PRD-001",
					ProductName:     "keyboard logitec",
					Price:           33,
					Volume:          94,
					Value:           310200000,
					TransactionDate: helper.ParseTime("2/23/22 22:08"),
					EntryDate:       helper.ParseTime("2/24/22 12:03"),
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 0 {
				result := tt.r.GetDatamar2(context.Background(), tx, 0)
				assert.Equal(t, tt.expected, result)
				return
			}
			if i == 1 {
				createDataDummy()
			}

			result := tt.r.GetDatamar2(context.Background(), tx, 0)

			datamart := result[0]
			datamart.TransactionDate = helper.ParseTime(helper.ParseTimeToString(datamart.TransactionDate))
			datamart.EntryDate = helper.ParseTime(helper.ParseTimeToString(datamart.EntryDate))
			result[0] = datamart
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_repositoryDatamartImpl_GetDatamar3(t *testing.T) {
	db := database.NewDB()
	tx, err := db.Begin()
	helper.PanicIFError(err)

	// this function for handle panic helper.CommitOrRollback
	defer func() {
		recover()
	}()
	defer helper.CommitOrRollback(tx)
	database.TruncateAllTable(db)

	repository := NewRepositoryDatamart()

	tests := []struct {
		name     string
		r        RepositoryDatamart
		want     int
		expected []domain.Datamart3
		err      string
	}{
		{
			name:     "get datamart null data",
			r:        repository,
			want:     0,
			expected: []domain.Datamart3(nil),
		},
		{
			name: "get datamart success",
			r:    repository,
			want: 0,
			expected: []domain.Datamart3{
				{
					CustomerId:      "CTM-001",
					CustomerName:    "hendrik",
					Price:           33,
					Volume:          94,
					Value:           310200000,
					TransactionDate: helper.ParseTime("2/23/22 22:08"),
					EntryDate:       helper.ParseTime("2/24/22 12:03"),
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 0 {
				result := tt.r.GetDatamar3(context.Background(), tx, 0)
				assert.Equal(t, tt.expected, result)
				return
			}
			if i == 1 {
				createDataDummy()
			}

			result := tt.r.GetDatamar3(context.Background(), tx, 0)

			datamart := result[0]
			datamart.TransactionDate = helper.ParseTime(helper.ParseTimeToString(datamart.TransactionDate))
			datamart.EntryDate = helper.ParseTime(helper.ParseTimeToString(datamart.EntryDate))
			result[0] = datamart
			assert.Equal(t, tt.expected, result)
		})
	}
}
