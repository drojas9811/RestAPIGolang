package helpers

import (
	"RestAPIGolang/internal/api/handlers"
	"RestAPIGolang/internal/database"
	model "RestAPIGolang/internal/models"
	"fmt"
	"log"
)

func seedAccount(fname, lname, pw string) *model.Account {
	acc, err := handlers.NewAccount(fname, lname, pw)
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
