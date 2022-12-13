package service_product

import (
	"context"
	"database/sql"
	"errors"
	"moonlay/database"
	repository_product "moonlay/repository/product"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newServiceProduct() (ServiceProduct, *sql.DB) {
	repoProduct := repository_product.NewRepositoryProduct()
	db := database.NewDB()
	return NewServiceProduct(repoProduct, db), db
}

func TestNewServiceProduct(t *testing.T) {
	repoProduct := repository_product.NewRepositoryProduct()
	db := database.NewDB()
	service := NewServiceProduct(repoProduct, db)

	reflection := reflect.TypeOf(service)

	assert.Equal(t, "*service_product.serviceProductImpl", reflection.String())
}

func Test_serviceProductImpl_CreateProductFromFile(t *testing.T) {
	serviceProduct, db := newServiceProduct()
	database.TruncateAllTable(db)
	// tx, err := db.Begin()
	// helper.PanicIFError(err)

	tests := []struct {
		name     string
		s        ServiceProduct
		want     [][]string
		expected error
	}{
		{
			name: "save product success",
			s:    serviceProduct,
			want: [][]string{
				{
					"test",
					"asdkfjhdask",
					"lasdkhfklsf",
					"kilsdjuksdjfhkl",
				},
				{
					"1",
					"PRD-001",
					"keyboard",
					"LRD",
					"Y",
				},
			},
			expected: errors.New(""),
		},
		{
			name: "save product failed",
			s:    serviceProduct,
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
			name: "save product failed close connection",
			s:    serviceProduct,
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
					tt.s.CreateProductFromFile(context.Background(), tt.want)
				})
			}
			if i == 1 {
				assert.PanicsWithError(t, tt.expected.Error(), func() {
					tt.s.CreateProductFromFile(context.Background(), tt.want)
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
				tt.s.CreateProductFromFile(context.Background(), tt.want)
			}
		})
	}
}
