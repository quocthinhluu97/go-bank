package useraccounts

import (
	"github.com/quocthinhluu97/go-bank/helpers"
	"github.com/quocthinhluu97/go-bank/interfaces"
	"fmt"
)


// func updateAccount(id uint, amount int) {
// 	db := helpers.ConnectDB()
// 	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)

// 	defer db.Close()
// }

func getAccount(id uint) *interfaces.Account {
	db := helpers.ConnectDB()
	account := &interfaces.Account{}

	if db.Where("id = ?", id).First(&account).RecordNotFound() {
		return nil
	}
	defer db.Close()
	return account
}

func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
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

		if int(fromAccount.Balance) < amount {
			return map[string]interface{}{"Message": "Accouont balance is not enough"}
		}

		updatedAccount := updateAccount(from, int(fromAccount.Balance) - amount)
		updateAccount(to, int(fromAccount.Balance) + amount)

		var response = map[string]interface{}{"Message": "all is fine"}
		response["data"] = updatedAccount
		return response
	}

	return map[string]interface{}{"Message": "not valid token"}
}


func updateAccount(id uint, amount int) interfaces.ResponseAccount {
	db := helpers.ConnectDB()
	account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}

	db.Where("id = ?", id).First(&account)
	account.Balance = uint(amount)
	db.Save(&account)


	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = int(account.Balance)

	defer db.Close()

	return responseAcc
}
