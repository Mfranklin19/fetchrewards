package main

import (
	"net/http"

	"fetchrewards.com/points/balance"
	"fetchrewards.com/points/expense"
	"fetchrewards.com/points/payer"
	"fetchrewards.com/points/transaction"
)

const apiBasePath = "/api"

func main() {
	transaction.SetupRoutes(apiBasePath)
	payer.SetupRoutes(apiBasePath)
	balance.SetupRoutes(apiBasePath)
	expense.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
