package domain

import (
	"reflect"
	"time"
)

type Datamart1 struct {
	TransactionId        string    `excel:"transaction_id"`
	BuyerId              string    `excel:"buyer_id"`
	BuyerName            string    `excel:"buyer_name"`
	SellerId             string    `excel:"seller_id"`
	SellerName           string    `excel:"seller_name"`
	ProductId            string    `excel:"product_id"`
	ProductName          string    `excel:"product_name"`
	Currency             string    `excel:"currency"`
	Price                int       `excel:"price"`
	Volume               int       `excel:"volume"`
	Value                int64     `excel:"value"`
	TransactionDate      time.Time `excel:"transaction_date"`
	TransactionMonth     int       `excel:"transaction_month"`
	TransactionYear      int       `excel:"transaction_year"`
	EntryDate            time.Time `excel:"entry_date"`
	EntryMonth           int       `excel:"entry_month"`
	EntryYear            int       `excel:"entry_year"`
	Buysell              string    `excel:"buy_sell"`
	ConfirmStatus        string    `excel:"confirm_status"`
	CompleteStatusBuyer  string    `excel:"complete_status_buyer"`
	CompleteStatusSeller string    `excel:"complete_status_seller"`
}

type Datamart3 struct {
	CustomerId      string    `excel:"customer_id"`
	CustomerName    string    `excel:"customer_name"`
	Price           int       `excel:"price"`
	Volume          int       `excel:"volume"`
	Value           int       `excel:"value"`
	TransactionDate time.Time `excel:"transaction_date"`
	EntryDate       time.Time `excel:"entry_date"`
}
type Datamart2 struct {
	ProductId       string    `excel:"product_id"`
	ProductName     string    `excel:"product_name"`
	Price           int       `excel:"price"`
	Volume          int       `excel:"volume"`
	Value           int       `excel:"value"`
	TransactionDate time.Time `excel:"transaction_date"`
	EntryDate       time.Time `excel:"entry_date"`
}

func (datamart *Datamart1) GetNumField(data Datamart1) []string {
	rf := reflect.TypeOf(data)
	var result []string

	for i := rf.NumField(); i > 0; i-- {
		result = append(result, rf.Field(i).Name)
	}
	return result
}

func (datamart *Datamart1) ChangeReflectValueToDataMart1(data reflect.Value) *Datamart1 {
	return &Datamart1{
		TransactionId:        data.FieldByName("TransactionId").String(),
		BuyerId:              data.FieldByName("BuyerId").String(),
		BuyerName:            data.FieldByName("BuyerName").String(),
		SellerId:             data.FieldByName("SellerId").String(),
		SellerName:           data.FieldByName("SellerName").String(),
		ProductId:            data.FieldByName("ProductId").String(),
		ProductName:          data.FieldByName("ProductName").String(),
		Currency:             data.FieldByName("Currency").String(),
		Price:                data.FieldByName("Price").Interface().(int),
		Volume:               data.FieldByName("Volume").Interface().(int),
		Value:                data.FieldByName("Value").Int(),
		TransactionDate:      data.FieldByName("TransactionDate").Interface().(time.Time),
		TransactionMonth:     data.FieldByName("TransactionMonth").Interface().(int),
		TransactionYear:      data.FieldByName("TransactionYear").Interface().(int),
		EntryDate:            data.FieldByName("EntryDate").Interface().(time.Time),
		EntryMonth:           data.FieldByName("EntryMonth").Interface().(int),
		EntryYear:            data.FieldByName("EntryYear").Interface().(int),
		Buysell:              data.FieldByName("Buysell").String(),
		ConfirmStatus:        data.FieldByName("ConfirmStatus").String(),
		CompleteStatusBuyer:  data.FieldByName("CompleteStatusBuyer").String(),
		CompleteStatusSeller: data.FieldByName("CompleteStatusSeller").String(),
	}
}
