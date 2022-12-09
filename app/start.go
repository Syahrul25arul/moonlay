package app

import (
	"context"
	"fmt"
	"moonlay/database"
	"moonlay/excel"
	"moonlay/helper"
	"time"

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
	start := time.Now()

	// TODO: Get file excel
	exc := excel.NewExcel()
	file := exc.GetFile("datamart.xlsx")

	// * set database connection and close at the end

	sql := database.NewDB()
	defer sql.Close()
	fmt.Println(fmt.Sprintf("duration execution setup app %s", time.Since(start)))

	timeReadDataFromExcel := time.Now()
	// TODO: get data customers from file excel

	// customers, err := helper.ReadData("customers")
	customers, err := file.ReadData("customers", "dummies-migration-tests.xlsx")
	helper.PanicIFError(err)

	// TODO: get data products from file excel

	products, err := file.ReadData("products", "dummies-migration-tests.xlsx")
	helper.PanicIFError(err)

	// TODO: get data transactions from file excel
	transactions, err := file.ReadData("transactions", "dummies-migration-tests.xlsx")
	helper.PanicIFError(err)
	fmt.Println(fmt.Sprintf("Duration execution for read data from excel %s", time.Since(timeReadDataFromExcel)))

	// * create repo and service customer
	repositoryCustomer := repository_customer.NewRepositoryCustomer()
	serviceCustomer := service_customer.NewServiceCustomer(repositoryCustomer, sql)

	// * create repo and service product
	repositoryProducts := repository_product.NewRepositoryProduct()
	serviceProduct := service_product.NewServiceProduct(repositoryProducts, sql)

	// * create repo and service transaction
	repositoryTransaction := repository_transaction.NewRepositoryTransaction()
	servicerTransaction := service_transaction.NewServiceTransaction(repositoryTransaction, sql)

	// * create repo and service datamart
	repositoryDatamart := repository_datamart.NewRepositoryDatamart()
	serviceDatamart := service_datamart.NewServiceDatamart(repositoryDatamart, sql)

	timeCreateDataToDB := time.Now()
	// TODO: save data customer from excel to db
	serviceCustomer.CreateCustomerFromFile(context.Background(), customers)

	// TODO: save data product from excel to db
	serviceProduct.CreateProductFromFile(context.Background(), products)

	// TODO: save data transactions from excel to db
	servicerTransaction.CreateTransactionFromFile(context.Background(), transactions)
	fmt.Println(fmt.Sprintf("duration execution save data to database %s", time.Since(timeCreateDataToDB)))

	totalDatamart1 := serviceDatamart.GetTotalData(context.Background())

	timeCreateDatamar1ToExcel := time.Now()
	// TODO: scheduler task for set datamart1 to file excel
	scheduler.StartScheduler(totalDatamart1, context.Background(), "Datamart1", "datamart", func(i int, sheetName string, domain string) {
		data := serviceDatamart.GetDatamar1(context.Background(), i)
		activeSell := int32((i * 100) + 1)
		file.CreateDataMart1(sheetName, data, activeSell, i)
	})
	fmt.Println(fmt.Sprintf("Duration execution for create datamart1 to excel %s", time.Since(timeCreateDatamar1ToExcel)))

	timeCreateDatamar2ToExcel := time.Now()
	// TODO: scheduler task for set datamart2 to file excel
	scheduler.StartScheduler(totalDatamart1, context.Background(), "Datamart2", "datamart2", func(i int, sheetName string, domain string) {
		data := serviceDatamart.GetDatamar2(context.Background(), i)
		activeSell := int32((i * 100) + 1)
		file.CreateDataMart2(sheetName, data, activeSell, i)
	})
	fmt.Println(fmt.Sprintf("Duration execution for create datamart2 to excel %s", time.Since(timeCreateDatamar2ToExcel)))

	timeCreateDatamar3ToExcel := time.Now()
	// TODO: scheduler task for set datamart3 to file excel
	scheduler.StartScheduler(totalDatamart1, context.Background(), "Datamart3", "datamart3", func(i int, sheetName string, domain string) {
		data := serviceDatamart.GetDatamar3(context.Background(), i)
		activeSell := int32((i * 100) + 1)
		file.CreateDataMart3(sheetName, data, activeSell, i)
	})
	fmt.Println(fmt.Sprintf("Duration execution for create datamart3 to excel %s", time.Since(timeCreateDatamar3ToExcel)))
	timeBeforeFinish := time.Now()
	fmt.Println(fmt.Sprintf("Duration execution before finish to excel %s", time.Since(timeBeforeFinish)))
	fmt.Println(fmt.Sprintf("Duration executution %s", time.Since(start)))
}
