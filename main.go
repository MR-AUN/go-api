package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	addbook := addressBook{
		Firstname: "kiattisak",
		Lastname:  "janjam",
		Code:      1998,
		Phone:     "0855089934",
	}
	json.NewEncoder(w).Encode(addbook)
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku" + port)
	}

	return ":" + port
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!!!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getAddress", getAddressBookAll)
	http.ListenAndServe(getPort(), nil)
}

func main() {
	handleRequest()
}
