package main

import (
	"github.com/quocthinhluu97/go-bank/api"
)


func main() {
	api
}


func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}

}
