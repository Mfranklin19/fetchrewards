package balance

import (
	"sort"
	"sync"
	"time"

	"fetchrewards.com/points/payer"
)

var balanceMap = struct {
	sync.RWMutex
	m map[int]Balance
}{m: make(map[int]Balance)}

// Determine if user has enough points to make a payment
func PaymentPossble(points int) bool {
	return points <= getTotalBalance()
}

func MakePayment(points int) []Report {
	balances := getBalancesByDate()
	var balance Balance
	paymentMap := struct {
		m map[string]Report
	}{m: make(map[string]Report)}
	for points > 0 {
		balance, balances = balances[0], balances[1:]
		report := &Report{
			Payer: balance.Payer,
			Amount: 0,
		}
		if balance.Points > points {
			payer.AddOrUpdatePayer(balance.Payer, -balance.Points)
			balance.Points -= points
			report.Amount -= points
			updateBalance(balance)
		} else {
			payer.AddOrUpdatePayer(balance.Payer, -balance.Points)
			report.Amount -= balance.Points
			removeBalance(balance)
		}
		points -= balance.Points
		if _,ok := paymentMap.m[report.Payer]; ok {
			report.Amount += paymentMap.m[report.Payer].Amount
			paymentMap.m[report.Payer] = *report
		} else {
			paymentMap.m[report.Payer] = *report
		}
	}
	paymentReport := make([]Report, 0, len(paymentMap.m))
	for _, value := range paymentMap.m {
		paymentReport = append(paymentReport, value)
	}
	return paymentReport
}

func getTotalBalance() int {
	balanceMap.RLock()
	balancePoints := 0
	for balanceId := range balanceMap.m {
		balancePoints += balanceMap.m[balanceId].Points
	}
	balanceMap.RUnlock()
	return balancePoints
}

func getBalancesByDate() []Balance {
	balanceMap.RLock()
	balances := make([]Balance, 0, len(balanceMap.m))
	for _, value := range balanceMap.m {
		balances = append(balances, value)
	}
	balanceMap.RUnlock()
	sort.Slice(balances, func(i, j int) bool {
		return balances[i].Timestamp.Before(balances[j].Timestamp)
	})
	return balances
}

func updateBalance(balance Balance) {	
	balanceMap.Lock()
	balanceMap.m[balance.BalanceId] = balance
	balanceMap.Unlock()
}

func removeBalance(balance Balance) {
	balanceMap.Lock()
	delete(balanceMap.m, balance.BalanceId)
	balanceMap.Unlock()
}

func AddBalance(balanceId int, payer string, points int, Timestamp time.Time) {
	balance := &Balance {
		BalanceId: balanceId,
		Payer: payer,
		Points: points,
		Timestamp: Timestamp,
	}
	balanceMap.Lock()
	balanceMap.m[balance.BalanceId] = *balance
	balanceMap.Unlock()
}