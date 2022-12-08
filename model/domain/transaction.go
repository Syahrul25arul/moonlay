package domain

import (
	"moonlay/helper"
	"time"
)

type Transactions struct {
	Id                   int
	TransactionId        string
	TransactionUuid      string
	RelUuid              string
	BuyerId              string
	SellerId             string
	ProductId            string
	Price                int
	Volume               int
	Value                uint64
	TransactionDate      time.Time
	EntryDate            time.Time
	ConfirmDate          time.Time
	CompleteDataBuyer    time.Time
	CompleteDataSeller   time.Time
	BuySell              string
	IsAmmend             string
	IsCancel             string
	ConfirmStatus        string
	CompleteStatusBuyer  string
	CompleteStatusSeller string
	Status               string
}

func (p *Transactions) ConvertDataStringToTransactions(data []string) {
	p.TransactionId = data[1]
	p.TransactionUuid = data[2]
	p.RelUuid = data[3]
	p.BuyerId = data[4]
	p.SellerId = data[5]
	p.ProductId = data[6]
	p.Price = helper.ConvertFromString(data[7], "int").(int)
	p.Volume = helper.ConvertFromString(data[8], "int").(int)
	p.Value = helper.ConvertFromString(data[9], "uint64").(uint64)
	p.TransactionDate = helper.ConvertFromString(data[10], "time").(time.Time)
	p.EntryDate = helper.ConvertFromString(data[11], "time").(time.Time)
	p.ConfirmDate = helper.ConvertFromString(data[12], "time").(time.Time)
	p.CompleteDataBuyer = helper.ConvertFromString(data[13], "time").(time.Time)
	p.CompleteDataSeller = helper.ConvertFromString(data[14], "time").(time.Time)
	p.BuySell = data[15]
	p.IsAmmend = data[16]
	p.IsCancel = data[17]
	p.ConfirmStatus = data[18]
	p.CompleteStatusBuyer = data[19]
	p.CompleteStatusSeller = data[20]
	p.Status = data[21]

}
