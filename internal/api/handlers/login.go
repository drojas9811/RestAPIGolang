package handlers

import (
	"RestAPIGolang/internal/auth"
	"RestAPIGolang/internal/database"
	model "RestAPIGolang/internal/models"
	"RestAPIGolang/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) error {
	var req model.LoginRequest 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	acc, err := database.GetAccountByNumber(int(req.Number))
	if err != nil {
		return err
	}

	if !acc.ValidPassword(req.Password) {
		return fmt.Errorf("not authenticated")
	}

	token, err := auth.CreateJWT(acc)
	if err != nil {
		return err
	}

	resp := model.LoginResponse{ //add  model
		Token:  token,
		Number: acc.Number,
	}

	return utils.WriteJSON(w, http.StatusOK, resp)
}
