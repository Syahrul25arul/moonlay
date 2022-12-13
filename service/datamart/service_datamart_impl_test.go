package service_datamart

import (
	"context"
	"database/sql"
	"errors"
	"moonlay/database"
	"moonlay/helper"
	"moonlay/model/domain"
	repository_customer "moonlay/repository/customer"
	repository_datamart "moonlay/repository/datamart"
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
	helper.CommitOrRollback(tx)
}

func newServiceDatamart() (ServiceDatamart, *sql.DB) {
	repoCustomer := repository_datamart.NewRepositoryDatamart()
	db := database.NewDB()
	return NewServiceDatamart(repoCustomer, db), db
}

func TestNewServiceDatamart(t *testing.T) {
	repoCustomer := repository_datamart.NewRepositoryDatamart()
	db := database.NewDB()
	service := NewServiceDatamart(repoCustomer, db)

	reflection := reflect.TypeOf(service)

	assert.Equal(t, "*service_datamart.serviceDatamartImpl", reflection.String())
}

func Test_serviceDatamartImpl_GetDatamar1(t *testing.T) {
	serviceDatamart, db := newServiceDatamart()
	tx, err := db.Begin()
	helper.PanicIFError(err)
	database.TruncateAllTable(db)

	createDataDummy()
	defer helper.CommitOrRollback(tx)

	tests := []struct {
		name     string
		s        ServiceDatamart
		want     int
		expected []domain.Datamart1
		err      error
	}{
		{
			name: "get datamart success",
			s:    serviceDatamart,
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
		{
			name:     "get datamart1 null data",
			s:        serviceDatamart,
			want:     0,
			expected: []domain.Datamart1(nil),
		},
		{
			name:     "get datamart1 close connection",
			s:        serviceDatamart,
			want:     0,
			expected: []domain.Datamart1(nil),
			err:      errors.New("sql: database is closed"),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 1 {
				database.TruncateAllTable(db)
				result := tt.s.GetDatamar1(context.Background(), tt.want)
				assert.Equal(t, tt.expected, result)
				return
			}

			if i == 2 {
				defer func() {
					if r := recover(); r != nil {
						errs := r.(error)
						assert.Equal(t, tt.err.Error(), errs.Error())
					}
				}()
				db.Close()
				tt.s.GetDatamar1(context.Background(), tt.want)
				return
			}

			result := tt.s.GetDatamar1(context.Background(), tt.want)
			datamart := result[0]
			datamart.TransactionDate = helper.ParseTime(helper.ParseTimeToString(datamart.TransactionDate))
			datamart.EntryDate = helper.ParseTime(helper.ParseTimeToString(datamart.EntryDate))
			result[0] = datamart
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_serviceDatamartImpl_GetTotalData(t *testing.T) {
	serviceDatamart, db := newServiceDatamart()
	tx, err := db.Begin()
	helper.PanicIFError(err)
	database.TruncateAllTable(db)

	createDataDummy()
	defer helper.CommitOrRollback(tx)

	tests := []struct {
		name     string
		s        ServiceDatamart
		expected int
		err      error
	}{
		{
			name:     "get total data success",
			s:        serviceDatamart,
			expected: 1,
		},
		{
			name:     "get total null",
			s:        serviceDatamart,
			expected: 0,
		},
		{
			name:     "get total data failed",
			s:        serviceDatamart,
			expected: 1,
			err:      errors.New("sql: database is closed"),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 1 {
				database.TruncateAllTable(db)
			}
			if i == 2 {
				db.Close()
				defer func() {
					assert.PanicsWithError(t, tt.err.Error(), func() {
						tt.s.GetTotalData(context.Background())
					})
				}()
				return
			}
			result := tt.s.GetTotalData(context.Background())
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_serviceDatamartImpl_GetDatamar2(t *testing.T) {
	serviceDatamart, db := newServiceDatamart()
	tx, err := db.Begin()
	helper.PanicIFError(err)
	database.TruncateAllTable(db)

	createDataDummy()
	defer helper.CommitOrRollback(tx)

	tests := []struct {
		name     string
		s        ServiceDatamart
		want     int
		expected []domain.Datamart2
		err      error
	}{
		{
			name: "get datamart2 success",
			s:    serviceDatamart,
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
		{
			name:     "get datamart2 null data",
			s:        serviceDatamart,
			want:     0,
			expected: []domain.Datamart2(nil),
		},
		{
			name:     "get datamart2 close connection",
			s:        serviceDatamart,
			want:     0,
			expected: []domain.Datamart2(nil),
			err:      errors.New("sql: database is closed"),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 1 {
				database.TruncateAllTable(db)
				result := tt.s.GetDatamar2(context.Background(), tt.want)
				assert.Equal(t, tt.expected, result)
				return
			}

			if i == 2 {
				defer func() {
					if r := recover(); r != nil {
						errs := r.(error)
						assert.Equal(t, tt.err.Error(), errs.Error())
					}
				}()
				db.Close()
				tt.s.GetDatamar2(context.Background(), tt.want)
				return
			}

			result := tt.s.GetDatamar2(context.Background(), tt.want)
			datamart := result[0]
			datamart.TransactionDate = helper.ParseTime(helper.ParseTimeToString(datamart.TransactionDate))
			datamart.EntryDate = helper.ParseTime(helper.ParseTimeToString(datamart.EntryDate))
			result[0] = datamart
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_serviceDatamartImpl_GetDatamar3(t *testing.T) {
	serviceDatamart, db := newServiceDatamart()
	tx, err := db.Begin()
	helper.PanicIFError(err)
	database.TruncateAllTable(db)

	createDataDummy()
	defer helper.CommitOrRollback(tx)

	tests := []struct {
		name     string
		s        ServiceDatamart
		want     int
		expected []domain.Datamart3
		err      error
	}{
		{
			name: "get datamart3 success",
			s:    serviceDatamart,
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
		{
			name:     "get datamart3 null data",
			s:        serviceDatamart,
			want:     0,
			expected: []domain.Datamart3(nil),
		},
		{
			name:     "get datamart3 close connection",
			s:        serviceDatamart,
			want:     0,
			expected: []domain.Datamart3(nil),
			err:      errors.New("sql: database is closed"),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 1 {
				database.TruncateAllTable(db)
				result := tt.s.GetDatamar3(context.Background(), tt.want)
				assert.Equal(t, tt.expected, result)
				return
			}

			if i == 2 {
				defer func() {
					if r := recover(); r != nil {
						errs := r.(error)
						assert.Equal(t, tt.err.Error(), errs.Error())
					}
				}()
				db.Close()
				tt.s.GetDatamar3(context.Background(), tt.want)
				return
			}

			result := tt.s.GetDatamar3(context.Background(), tt.want)
			datamart := result[0]
			datamart.TransactionDate = helper.ParseTime(helper.ParseTimeToString(datamart.TransactionDate))
			datamart.EntryDate = helper.ParseTime(helper.ParseTimeToString(datamart.EntryDate))
			result[0] = datamart
			assert.Equal(t, tt.expected, result)
		})
	}
}
