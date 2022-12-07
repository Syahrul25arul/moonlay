package domain

import "time"

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
