package api

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/quocthinhluu97/go-bank/helpers"
	"github.com/quocthinhluu97/go-bank/users"
	"github.com/quocthinhluu97/go-bank/useraccounts"
	"github.com/quocthinhluu97/go-bank/transactions"
	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email string
	Password string
}

type TransactionBody struct {
	UserId uint
	From uint
	To uint
	Amount uint
}

func StartAPI() {
	router := mux.NewRouter()

	// Add panic handler middleware
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/transaction", transaction).Methods("POST")
	router.HandleFunc("/transactions/{userID}", getMyTransactions).Methods("GET")
	fmt.Println("App is working on port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody Register
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	apiResponse(login, w)
}

func register(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody Register
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	apiResponse(register, w)
}

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["Message"] == "all is fine" {
		resp := call
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := call
		json.NewEncoder(w).Encode(resp)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)
	apiResponse(user, w)

}

func transaction(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	auth := r.Header.Get("Authorization")
	var formattedBody TransactionBody
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	transaction := useraccounts.Transaction(formattedBody.UserId, formattedBody.From,formattedBody.To, formattedBody.Amount, auth)
	apiResponse(transaction, w)
}

func getMyTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userID"]
	auth := r.Header.Get("Authorization")

	transactions := transactions.GetMyTransactions(userId, auth)
	apiResponse(transactions, w)
}
