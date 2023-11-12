package handlers

import (
	"RestAPIGolang/internal/database"
	"RestAPIGolang/internal/helpers"
	"RestAPIGolang/internal/models"
	"RestAPIGolang/internal/utils"
	"encoding/json"
	"net/http"
)

func GetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := database.GetAccounts()
	if err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, accounts)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) error {
	req := new(models.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	account := new(models.Account)
	err := account.NewAccount(req.FirstName, req.LastName, req.Password)
	if err != nil {
		return err
	}
	if err := database.CreateAccount(account); err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, account)
}

func GetAccountByID(w http.ResponseWriter, r *http.Request) error {
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

func DeleteAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetID(r)
	if err != nil {
		return err
	}
	if err := database.DeleteAccount(id); err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}
