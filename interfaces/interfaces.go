package interfaces

import (
	"github.com/jinzhu/gorm"
)

type ResponseTransaction struct {
	Id uint
	From uint
	To uint
	Amount int
}

type User struct {
	gorm.Model
	Username string
	Email string
	Password string
}

type Account struct {
	gorm.Model
	Type string
	Username string
	Balance uint
	UserID uint
}

type ResponseAccount struct {
	ID  uint
	Balance uint
	Username string
 }

type ResponseUser struct {
	ID  uint
	Username string
	Email string
	Accounts []ResponseAccount

}

type Validation struct {
	Value string
	Valid string
}

type ErrResponse struct {
	Message string
}

type Transaction struct {
	gorm.Model
	From uint
	To uint
	Amount uint
}
