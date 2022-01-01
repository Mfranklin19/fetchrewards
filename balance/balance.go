package balance

import (
	"time"
)

type Balance struct {
	BalanceId int
	Payer string
	Points int
	Timestamp time.Time
}