package transactions

import (
	"github.com/quocthinhluu97/go-bank/helpers"
	"github.com/quocthinhluu97/go-bank/database"
	"github.com/quocthinhluu97/go-bank/interfaces"
)



func CreateTransaction(From uint, To uint, Amount uint) {
	transaction := &interfaces.Transaction{From: From, To: To, Amount: Amount}
	database.DB.Create(&transaction)

}


func GetTransactionsByAccount(id uint) []interfaces.ResponseTransaction {
	transactions := []interfaces.ResponseTransaction{}
	database.DB.Table("transactions").Select(
		"id, transactions.from, transactions.to, amount").Where(
		interfaces.Transaction{From: id}).Or(
		interfaces.Transaction{To: id}).Scan(&transactions)

	return transactions
}

func GetMyTransactions(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)
	if isValid {
		accounts := []interfaces.ResponseAccount{}
		database.DB.Table("accounts").Select(
			"id, name, balance").Where(
			"user_id = ?", id).Scan(&accounts)

		transactions := []interfaces.ResponseTransaction{}

		for i := 0; i < len(accounts); i++ {
			accTransactions := GetTransactionsByAccount(accounts[i].ID)
			transactions = append(transactions, accTransactions...)
		}

		var response = map[string]interface{}{"Message": "all is fine"}
		response["data"] = transactions
		return response

	} else {
		return map[string]interface{}{"Message": "Not Valid Token"}
	}
}
