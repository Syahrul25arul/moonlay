package service_transaction

import (
	"context"
	"database/sql"
	"errors"
	"moonlay/database"
	repositor_transaction "moonlay/repository/transactions"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newServiceTransaction() (ServiceTransaction, *sql.DB) {
	repoTransaction := repositor_transaction.NewRepositoryTransaction()
	db := database.NewDB()
	return NewServiceTransaction(repoTransaction, db), db
}
func TestNewServiceTransaction(t *testing.T) {
	repoTransaction := repositor_transaction.NewRepositoryTransaction()
	db := database.NewDB()
	service := NewServiceTransaction(repoTransaction, db)

	reflection := reflect.TypeOf(service)

	assert.Equal(t, "*service_transaction.serviceTransactionImpl", reflection.String())
}

func Test_serviceTransactionImpl_CreateTransactionFromFile(t *testing.T) {
	serviceTransaction, db := newServiceTransaction()
	database.TruncateAllTable(db)
	// tx, err := db.Begin()
	// helper.PanicIFError(err)

	tests := []struct {
		name     string
		s        ServiceTransaction
		want     [][]string
		expected error
	}{
		{
			name: "save transaction success",
			s:    serviceTransaction,
			want: [][]string{
				{
					"test",
					"asdkfjhdask",
					"lasdkhfklsf",
					"kilsdjuksdjfhkl",
				},
				{
					"1",
					"TR00-1",
					"10fc653e-9f40",
					"026a76a2-7bab",
					"CTM00-89",
					"CTM00-130",
					"PRD00-894",
					"33",
					"94",
					"310200000",
					"2/23/22 22:08",
					"2/24/22 12:03",
					"2/25/22 12:10",
					"2/26/22 15:50",
					"2/27/22 12:23",
					"B",
					"N",
					"N",
					"CONFIRMED",
					"COMPLETE1",
					"COMPLETE1",
					"M",
				},
			},
			expected: errors.New(""),
		},
		{
			name: "save transaction failed",
			s:    serviceTransaction,
			want: [][]string{
				{},
				{
					"lsadnkf",
					"alskdfnm",
				},
			},
			expected: errors.New("runtime error: index out of range [2] with length 2"),
		},
		{
			name: "save transactiokn failed close connection",
			s:    serviceTransaction,
			want: [][]string{
				{},
				{
					"lsadnkf",
					"alskdfnm",
				},
			},
			expected: errors.New("sql: database is closed"),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 0 {
				assert.NotPanics(t, func() {
					tt.s.CreateTransactionFromFile(context.Background(), tt.want)
				})
			}
			if i == 1 {
				assert.PanicsWithError(t, tt.expected.Error(), func() {
					tt.s.CreateTransactionFromFile(context.Background(), tt.want)
				})
			}
			if i == 2 {
				db.Close()
				defer func() {
					if r := recover(); r != nil {
						errs := r.(error)
						assert.Equal(t, tt.expected.Error(), errs.Error())
					}
				}()
				tt.s.CreateTransactionFromFile(context.Background(), tt.want)
			}
		})
	}
}
