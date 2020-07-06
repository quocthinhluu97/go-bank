package main

import (
	"github.com/quocthinhluu97/go-bank/api"
	"github.com/quocthinhluu97/go-bank/database"
	// "github.com/quocthinhluu97/go-bank/migrations"
)

func main() {
	database.InitDatabase()
	api.StartAPI()
	// migrations.MigrateTransactions()
	// migrations.Migrate()
}
