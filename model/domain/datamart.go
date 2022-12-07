package domain

import (
	"reflect"
	"time"
)

type Datamart1 struct {
	TransactionId        string
	BuyerId              string
	BuyerName            string
	SellerId             string
	SellerName           string
	ProductId            string
	ProductName          string
	Currency             string
	Price                int
	Volume               int
	Value                int64
	TransactionDate      time.Time
	TransactionMonth     int
	TransactionYear      int
	EntryDate            time.Time
	EntryMonth           int
	EntryYear            int
	Buysell              string
	ConfirmStatus        string
	CompleteStatusBuyer  string
	CompleteStatusSeller string
}

type Datamart3 struct {
	CustomerId      string
	CustomerName    string
	Price           int
	Volume          int
	Value           int
	TransactionDate time.Time
	EntryDate       time.Time
}
type Datamart2 struct {
	ProductId       string
	ProductName     string
	Price           int
	Volume          int
	Value           int
	TransactionDate time.Time
	EntryDate       time.Time
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
