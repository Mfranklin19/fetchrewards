package main

import (
	"net/http"

	"fetchrewards.com/points/transaction"
)

const apiBasePath = "/api"

func main() {
	transaction.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
