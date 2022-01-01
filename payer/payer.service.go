package payer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fetchrewards.com/points/cors"
)

const apiTransactionsPath = "payers"

func SetupRoutes(apiBasePath string) {
	handlePayers := http.HandlerFunc(payersHandler)
	//handleTransaction := http.HandlerFunc(transactionHandler)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, apiTransactionsPath), cors.Middleware(handlePayers))
	//http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, apiTransactionsPath), cors.Middleware(handleTransaction))
}

func payersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		payers := getPayerList()
		payersJson, err := json.Marshal(payers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(payersJson)
	// Handled by middleware. Used for returning CORS headers
	case http.MethodOptions:
		return
	}
}