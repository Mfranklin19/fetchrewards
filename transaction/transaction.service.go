package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"fetchrewards.com/points/cors.go"
)

const apiTransactionsPath = "transactions"

func SetupRoutes(apiBasePath string) {
	handleTransactions := http.HandlerFunc(transactionsHandler)
	handleTransaction := http.HandlerFunc(transactionHandler)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, apiTransactionsPath), cors.Middleware(handleTransactions))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, apiTransactionsPath), cors.Middleware(handleTransaction))
}

func transactionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		transactions := getTransactionList()
		transactionsJson, err := json.Marshal(transactions)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(transactionsJson)
	case http.MethodPost:
		var newTransaction Transaction
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.Unmarshal(bodyBytes, &newTransaction)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newTransaction.TransactionId != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = addOrUpdateTransaction(newTransaction)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	// Handled by middleware. Used for returning CORS headers
	case http.MethodOptions:
		return
	}
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "transactions/")
	transactionId, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	transaction := getTransactionById(transactionId)
	if transaction == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		transactionJson, err := json.Marshal(transaction)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(transactionJson)
	case http.MethodPut:
		var updatedTransaction Transaction
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updatedTransaction)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedTransaction.TransactionId != transactionId {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = addOrUpdateTransaction(updatedTransaction)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodDelete:
		removeTransaction(transactionId)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
