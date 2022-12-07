package app

import (
	"context"
	"errors"
	"moonlay/helper"

	repository_customer "moonlay/repository/customer"
	repository_datamart "moonlay/repository/datamart"
	repository_product "moonlay/repository/product"
	repository_transaction "moonlay/repository/transactions"
	"moonlay/scheduler"
	service_customer "moonlay/service/customer"
	service_datamart "moonlay/service/datamart"
	service_product "moonlay/service/product"
	service_transaction "moonlay/service/transaction"
)

func Start() {
	helper.GetFileCsv("datamart.xlsx")
	f, fileExists := helper.CheckFile("datamart.xlsx")
	if fileExists {
		helper.PanicIFError(errors.New("file not exists"))
	}

	sql := NewDB()
	defer sql.Close()

	// customers, err := helper.ReadData("customers")
	customers, err := helper.ReadData("customers", "dummies-migration-tests.xlsx")
	helper.PanicIFError(err)

	products, err := helper.ReadData("products", "dummies-migration-tests.xlsx")
	helper.PanicIFError(err)

	transactions, err := helper.ReadData("transactions", "dummies-migration-tests.xlsx")
	helper.PanicIFError(err)

	repositoryCustomer := repository_customer.NewRepositoryCustomer()
	serviceCustomer := service_customer.NewServiceCustomer(repositoryCustomer, sql)

	repositoryProducts := repository_product.NewRepositoryProduct()
	serviceProduct := service_product.NewServiceProduct(repositoryProducts, sql)

	repositoryTransaction := repository_transaction.NewRepositoryTransaction()
	servicerTransaction := service_transaction.NewServiceTransaction(repositoryTransaction, sql)

	repositoryDatamart := repository_datamart.NewRepositoryDatamart()
	serviceDatamart := service_datamart.NewServiceDatamart(repositoryDatamart, sql)

	serviceCustomer.CreateCustomerFromFile(context.Background(), customers)
	serviceProduct.CreateProductFromFile(context.Background(), products)
	servicerTransaction.CreateTransactionFromFile(context.Background(), transactions)

	totalDatamart1 := serviceDatamart.GetTotalData(context.Background())

	scheduler.StartScheduler(totalDatamart1, context.Background(), f, "Datamart1", "datamart", func(i int, sheetName string, domain string) {
		data := serviceDatamart.GetDatamar1(context.Background(), i)
		activeSell := (i * 100) + 1
		helper.CreateDataMart1(f, sheetName, data, activeSell, i)
	})

	scheduler.StartScheduler(totalDatamart1, context.Background(), f, "Datamart2", "datamart2", func(i int, sheetName string, domain string) {
		data := serviceDatamart.GetDatamar2(context.Background(), i)
		activeSell := (i * 100) + 1
		helper.CreateDataMart2(f, sheetName, data, activeSell, i)
	})

	scheduler.StartScheduler(totalDatamart1, context.Background(), f, "Datamart3", "datamart3", func(i int, sheetName string, domain string) {
		data := serviceDatamart.GetDatamar3(context.Background(), i)
		activeSell := (i * 100) + 1
		helper.CreateDataMart3(f, sheetName, data, activeSell, i)
	})
}
