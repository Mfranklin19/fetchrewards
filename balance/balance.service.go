package balance

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fetchrewards.com/points/cors"
)

const apiBalancePath = "balance"

func SetupRoutes(apiBasePath string) {
	handleBalance := http.HandlerFunc(balanceHandler)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, apiBalancePath), cors.Middleware(handleBalance))
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		balance := getTotalBalance()
		balanceJSON, err := json.Marshal(balance)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(balanceJSON)
	}
}