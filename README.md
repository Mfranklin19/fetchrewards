# FetchRewards Points Project
---
## **Install**

### **Download Go**

If you do not already have Go installed on your machine, head to [the Go website](https://go.dev/doc/install) and follow the instructions to install Go for your operating system. We will use the Go CLI to run the project later.

### **Download Postman**

This is a headless program, so we will be interacting with it using third party software. You can use any tool you like to do so, but I prefer Postman. you can [download Postman here](https://www.postman.com/downloads/) if you choose to do the same!

## **Run the Project**

To run the project, first clone the repository to your machine using `git clone https://github.com/Mfranklin19/fetchrewards.git` with the command line tool of your choice. Then navigate to the project directory using `cd fetchrewards`. Finally, use the command `go run main.go`. The service will automatically start on your localhost at port 5000. I do not have it set up for you to configure port usage, so please ensure this port is available at runtime.

---

## **Modules**

### **Transaction**

This module is associated with adding points to the user's account. It tracks these points with the **Transactions** struct. This data is immutable, and is for records keeping purposes. Calling the post method for this module triggers the creation of all the connected modules below.

**Endpoints**:
- GET http:/localhost/api/transactions
    - Returns a list of all transactions performed in the system
- POST http:/localhost/api/transactions
    - Inserts a new transaction into the system. Transaction is defined as a json object in the body
    - Example Body: `{ "TransactionId": 0, "payer": "payer7", "points": 25,    "timestamp": "2002-10-31T10:34:24.507332-06:00"}`
- GET http:/localhost/api/transaction/{transactionId}
    - Get information about a specific transaction

### **Balance**

This module is associated with keeping track of the balance of the user's account. This module contains the Balance struct, which is a mutable mirror of the Transaction struct. This is the struct that gets taken from when spending points.

**Endpoints**
- GET http:/localhost/api/balance
    - Returns the total remaining balance on a user's account

### **Payer**

This module is associated with keeping track of different payers on the user's account.

**Endpoints**
- GET http:/localhost/api/payer
    - Returns a list of all payers and their respective point balances

### **Expense**

This module is associated with spending the points a user has.

**Endpoints**
- POST http:/localhost/api/spend
    - Attempts to spend points from the account. The spend amount is defined as a json object in the body
    - Example Body: `{"Amount":1000}`