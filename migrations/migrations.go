package migrations

import (
	"github.com/quocthinhluu97/go-bank/helpers"
	"github.com/quocthinhluu97/go-bank/interfaces"
	"github.com/quocthinhluu97/go-bank/database"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createAccounts() {

	users := &[2]interfaces.User {
		{Username: "Martin", Email: "martin@fakemail.com"},
		{Username: "Michael", Email: "michael@fakemail.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		database.DB.Create(&user)

		account := &interfaces.Account{
			Type: "Daily Account",
			Username: string(users[i].Username),
			Balance: uint(1000 * int(i+1)),
			UserID: user.ID,
		}

		database.DB.Create(&account)
	}
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	Transactions := &interfaces.Transaction{}
	database.DB.AutoMigrate(&User, &Account, &Transactions)

	createAccounts()
}


