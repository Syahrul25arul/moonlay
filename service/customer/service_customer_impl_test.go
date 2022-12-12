package service_customer

import (
	"context"
	"database/sql"
	"errors"
	"moonlay/database"
	repository_customer "moonlay/repository/customer"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newServiceCustomer() (ServiceCustomer, *sql.DB) {
	repoCustomer := repository_customer.NewRepositoryCustomer()
	db := database.NewDB()
	return NewServiceCustomer(repoCustomer, db), db
}

func TestNewServiceCustomer(t *testing.T) {
	repoCustomer := repository_customer.NewRepositoryCustomer()
	db := database.NewDB()
	service := NewServiceCustomer(repoCustomer, db)

	reflection := reflect.TypeOf(service)

	assert.Equal(t, "*service_customer.ServiceCustomerImpl", reflection.String())
}

func TestServiceCustomerImpl_CreateCustomerFromFile(t *testing.T) {
	serviceCustomer, db := newServiceCustomer()
	database.TruncateAllTable(db)
	// tx, err := db.Begin()
	// helper.PanicIFError(err)

	tests := []struct {
		name     string
		s        ServiceCustomer
		want     [][]string
		expected error
	}{
		{
			name: "save customer success",
			s:    serviceCustomer,
			want: [][]string{
				{
					"test",
					"asdkfjhdask",
					"lasdkhfklsf",
				},
				{
					"1",
					"CTM-001",
					"hendrik",
					"Y",
				},
			},
			expected: errors.New(""),
		},
		{
			name: "save customer failed",
			s:    serviceCustomer,
			want: [][]string{
				{},
				{
					"lsadnkf",
					"alskdfnm",
				},
			},
			expected: errors.New("runtime error: index out of range [2] with length 2"),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 0 {
				assert.NotPanics(t, func() {
					tt.s.CreateCustomerFromFile(context.Background(), tt.want)
				})
			}
			if i == 1 {
				assert.PanicsWithError(t, tt.expected.Error(), func() {
					tt.s.CreateCustomerFromFile(context.Background(), tt.want)
				})
			}
		})
	}
}
