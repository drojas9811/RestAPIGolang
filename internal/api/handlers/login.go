package handlers

import (
	"RestAPIGolang/internal/auth"
	"RestAPIGolang/internal/database"
	"RestAPIGolang/internal/models"
	"RestAPIGolang/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) error {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("%s", "LOG_ERROR::LOGIN::MSG::Unknown login model.")
		return err
	}
	log.Printf("%s", "LOG_INFO::LOGIN::MSG::Login model is right.")
	acc, err := database.GetAccountByNumber(int(req.Number))
	if err != nil {
		log.Printf("%s", "LOG_ERROR::LOGIN::MSG::Invalid number.")
		return err
	}
	log.Printf("%s", "LOG_INFO::LOGIN::MSG::Number has been found.")
	if !acc.ValidPassword(req.Password) {
		log.Printf("%s", "LOG_ERROR::LOGIN::MSG::Invalid password.")
		return fmt.Errorf("not authenticated")
	}
	log.Printf("%s", "LOG_INFO::LOGIN::MSG::Password has been valited successfully.")
	token, err := auth.CreateJWT(int(acc.Number), (acc.FirstName + " " + acc.LastName))
	if err != nil {
		log.Printf("%s", "LOG_ERROR::LOGIN::MSG::An error occurred while creating a JWT token.")
		return err
	}
	log.Printf("%s", "LOG_INFO::LOGIN::MSG::JWT has been created successfully.")
	resp := models.LoginResponse{ //add  model
		Token:  token,
		Number: acc.Number,
	}
	log.Printf("%s", "LOG_INFO::LOGIN::MSG::Login has been successfully.")
	return utils.WriteJSON(w, http.StatusOK, resp)
}
