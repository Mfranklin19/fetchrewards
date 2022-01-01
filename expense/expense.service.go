package expense

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fetchrewards.com/points/balance"
	"fetchrewards.com/points/cors"
)

const apiExpensePath = "spend"

func SetupRoutes(apiBasePath string) {
	handleExpense := http.HandlerFunc(expenseHandler)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, apiExpensePath), cors.Middleware(handleExpense))
}

func expenseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var expense Expense
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.Unmarshal(bodyBytes, &expense)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !balance.PaymentPossble(expense.Amount) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		report := balance.MakePayment(expense.Amount)
		balanceJSON, err := json.Marshal(report)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(balanceJSON)
		return
	}
}