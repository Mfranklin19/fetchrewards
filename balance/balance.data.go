package balance

import (
	"fmt"
	"sort"
	"sync"
)

var payerMap = struct {
	sync.RWMutex
	m map[int]Payer
}{m: make(map[int]Payer)}

func getPayerById(payerId int) *Payer {
	payerMap.RLock()
	defer payerMap.RUnlock()
	if transaction, ok := payerMap.m[payerId]; ok {
		return &transaction
	}
	return nil
}

func getPayerIds() []int {
	payerMap.RLock()
	payerIds := []int{}
	for key := range payerMap.m {
		payerIds = append(payerIds, key)
	}
	payerMap.RUnlock()
	sort.Ints(payerIds)
	return payerIds
}

func getNextPayerId() int {
	payerIds := getPayerIds()
	if len(payerIds) == 0 {
		return 1
	}
	return payerIds[len(payerIds)-1] + 1
}

func addOrUpdatePayer(payer Payer) (int, error) {
	addOrUpdateId := -1

	if payer.PayerId > 0 {
		oldTransaction := getPayerById(payer.PayerId)
		if oldTransaction == nil {
			return 0, fmt.Errorf("transaction id [%d] does not exist", payer.PayerId)
		}
		addOrUpdateId = payer.PayerId
	} else {
		addOrUpdateId = getNextPayerId()
		payer.PayerId = addOrUpdateId
	}
	payerMap.Lock()
	payerMap.m[payer.PayerId] = payer
	payerMap.Unlock()
	return addOrUpdateId, nil
}
