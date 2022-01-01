package payer

import (
	"sync"
)

var payerMap = struct {
	sync.RWMutex
	m map[string]Payer
}{m: make(map[string]Payer)}


func getPayerNames() []string {
	payerMap.RLock()
	payerIds := []string{}
	for key := range payerMap.m {
		payerIds = append(payerIds, key)
	}
	payerMap.RUnlock()
	return payerIds
}

func getPayerList() []Payer {
	payerMap.RLock()
	payers := make([]Payer, 0, len(payerMap.m))
	for _, value := range payerMap.m {
		payers = append(payers, value)
	}
	payerMap.RUnlock()
	
	return payers
}

func getPayerByName(payerName string) *Payer {
	payerMap.RLock()
	defer payerMap.RUnlock()
	if payer, ok := payerMap.m[payerName]; ok {
		return &payer
	}
	return nil
}

func AddOrUpdatePayer(payerName string, points int) (Payer, error) {
	payer := getPayerByName(payerName)	
	if payer == nil {
		payer = &Payer{
			PayerName: payerName,
			Points: points,
		}
	} else {
		payer.Points += points
	}
	payerMap.Lock()
	payerMap.m[payerName] = *payer
	payerMap.Unlock()
	return *payer, nil
}