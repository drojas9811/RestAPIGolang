package helpers

import (
	"fmt"
	"log"

	"github.com/anthdm/gobank/internal/api/handlers"
	"github.com/anthdm/gobank/internal/database"
	model "github.com/anthdm/gobank/internal/models"
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
