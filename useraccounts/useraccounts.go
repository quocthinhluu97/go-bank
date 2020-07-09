package useraccounts

import (
	"github.com/quocthinhluu97/go-bank/helpers"
	"github.com/quocthinhluu97/go-bank/interfaces"
	"github.com/quocthinhluu97/go-bank/database"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}

func getAccount(id uint) *interfaces.Account {
	account := &interfaces.Account{}

	if database.DB.Where("id = ?", id).First(&account).RecordNotFound() {
		return nil
	}
	return account
}

func Transaction(userId uint, from uint, to uint, amount uint, jwt string) map[string]interface{} {
	userIdString := fmt.Sprint(userId)
	isValid := helpers.ValidateToken(userIdString, jwt)

	if isValid {
		fromAccount := getAccount(from)
		toAccount := getAccount(to)

		if fromAccount == nil || toAccount == nil {
			return map[string]interface{}{"Message": "Account not found"}
		}

		if fromAccount.UserID != userId {
			return map[string]interface{}{"Message": "You are not the owner of the account"}
		}

		if uint(fromAccount.Balance) < amount {
			return map[string]interface{}{"Message": "Account balance is not enough"}
		}

		updatedAccount := updateAccount(from, uint(fromAccount.Balance) - amount)
		updateAccount(to, uint(toAccount.Balance) + amount)

		var response = map[string]interface{}{"Message": "all is fine"}
		response["from"] = updatedAccount
		response["amount"] = amount
		return response
	}

	return map[string]interface{}{"Message": "not valid token"}
}


func updateAccount(id uint, amount uint) interfaces.ResponseAccount {
	account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}

	database.DB.Where("id = ?", id).First(&account)
	account.Balance = uint(amount)
	database.DB.Save(&account)

	responseAcc.ID = account.ID
	responseAcc.Username = account.Username
	responseAcc.Balance = uint(account.Balance)

	return responseAcc
}
