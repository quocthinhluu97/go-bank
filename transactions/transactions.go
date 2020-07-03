package transactions

import (
	"github.com/quocthinhluu97/go-bank/helpers"
	"github.com/quocthinhluu97/go-bank/interfaces"
)



func CreateTransaction(From uint, To uint, Amount int) {
	db := helpers.ConnectDB()
	transaction := &interfaces.Transaction{From: From, To: To, Amount: Amount}
	db.Create(&transaction)

	defer db.Close()

}
