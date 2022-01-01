package transaction

import (
	"fmt"
	"sort"
	"sync"

	"fetchrewards.com/points/balance"
	"fetchrewards.com/points/payer"
)

var transactionMap = struct {
	sync.RWMutex
	m map[int]Transaction
}{m: make(map[int]Transaction)}

func getTransactionById(transactionId int) *Transaction {
	transactionMap.RLock()
	defer transactionMap.RUnlock()
	if transaction, ok := transactionMap.m[transactionId]; ok {
		return &transaction
	}
	return nil
}

func getTransactionList() []Transaction {
	transactionMap.RLock()
	transactions := make([]Transaction, 0, len(transactionMap.m))
	for _, value := range transactionMap.m {
		transactions = append(transactions, value)
	}
	transactionMap.RUnlock()
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Timestamp.Before(transactions[j].Timestamp)
	})
	return transactions
}

func getTransactionIds() []int {
	transactionMap.RLock()
	transactionIds := []int{}
	for key := range transactionMap.m {
		transactionIds = append(transactionIds, key)
	}
	transactionMap.RUnlock()
	sort.Ints(transactionIds)
	return transactionIds
}

func getNextTransactionId() int {
	transactionIds := getTransactionIds()
	if len(transactionIds) == 0 {
		return 1
	}
	return transactionIds[len(transactionIds)-1] + 1
}

func addOrUpdateTransaction(transaction Transaction) (int, error) {
	addOrUpdateId := -1
	if transaction.TransactionId > 0 {
		oldTransaction := getTransactionById(transaction.TransactionId)
		if oldTransaction == nil {
			return 0, fmt.Errorf("transaction id [%d] does not exist", transaction.TransactionId)
		}
		addOrUpdateId = transaction.TransactionId
	} else {
		addOrUpdateId = getNextTransactionId()
		transaction.TransactionId = addOrUpdateId
	}
	payer.AddOrUpdatePayer(transaction.Payer, transaction.Points)
	balance.AddBalance(transaction.TransactionId, transaction.Payer, transaction.Points, transaction.Timestamp)
	transactionMap.Lock()
	transactionMap.m[transaction.TransactionId] = transaction
	transactionMap.Unlock()
	return addOrUpdateId, nil
}
