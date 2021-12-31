package transaction

import "time"

type Transaction struct {
	TransactionId int       `json:"transactionId"`
	Payer         string    `json:"payer"`
	Points        int       `json:"points"`
	Timestamp     time.Time `json:"timestamp"`
}
