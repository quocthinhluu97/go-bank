package main


func main() {

}


func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}

}
