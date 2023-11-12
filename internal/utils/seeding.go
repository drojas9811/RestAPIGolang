package utils

import (
	"RestAPIGolang/internal/database"
	model "RestAPIGolang/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}


func seedAccount(fname, lname, pw string) *model.Account {
	acc:=new(model.Account)
	err := acc.NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acc.Number)

	return acc
}

func SeedAccounts() {
	seedAccount("anthony", "GG", "hunter88888")
}

