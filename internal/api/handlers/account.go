package handlers

import (
	"RestAPIGolang/internal/database"
	"RestAPIGolang/internal/helpers"
	model "RestAPIGolang/internal/models"
	"RestAPIGolang/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := database.GetAccounts()
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, accounts)
}

func GetAccountByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := helpers.GetID(r)
		if err != nil {
			return err
		}

		account, err := database.GetAccountByID(id)
		if err != nil {
			return err
		}

		return utils.WriteJSON(w, http.StatusOK, account)
	}

	if r.Method == "DELETE" {
		return DeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) error {
	req := new(model.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	account:= new(model.Account)
	err := account.NewAccount(req.FirstName, req.LastName, req.Password)
	if err != nil {
		return err
	}
	if err := database.CreateAccount(account); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, account)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetID(r)
	if err != nil {
		return err
	}

	if err := database.DeleteAccount(id); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

