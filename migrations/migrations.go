package migrations

import (
	"github.com/quocthinhluu97/go-bank/helpers"
	"github.com/quocthinhluu97/go-bank/interfaces"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


// func connectDB() *gorm.DB {
// 	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=gobank")
// 	HandleErr(err)
// 	return db
// }

func createAccounts() {
	db := helpers.ConnectDB()

	users := [2]interfaces.User{
		{Username: "Martin", Email: "martin@fakemail.com"},
		{Username: "Michael", Email: "michael@fakemail.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"),
			Balance: uint(1000 * int(i+1)),
			UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User, &Account)
	defer db.Close()


	createAccounts()
}



