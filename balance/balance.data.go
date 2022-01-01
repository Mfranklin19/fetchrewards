package balance

import (
	"sort"
	"sync"
	"time"
)

var balanceMap = struct {
	sync.RWMutex
	m map[int]Balance
}{m: make(map[int]Balance)}

func PaymentPossble(points int) bool {
	return points <= getTotalBalance()
}

func MakePayment(points int) (int,error) {
	balances := getBalancesByDate()
	var balance Balance
	for points > 0 {
		balance, balances = balances[0], balances[1:]
		if balance.Points > points {
			balance.Points -= points
			updateBalance(balance)
		} else {
			removeBalance(balance)
		}
		points -= balance.Points
	}
	return getTotalBalance(), nil
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