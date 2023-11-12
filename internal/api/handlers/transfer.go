package handlers

import (
	model "RestAPIGolang/internal/models"
	"RestAPIGolang/internal/utils"
	"encoding/json"
	"net/http"
)

func Transfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := new(model.TransferRequest) //add model
	if err := json.NewDecoder(r.Body).Decode(transferReq); err != nil {
		return err
	}
	defer r.Body.Close()

	return utils.WriteJSON(w, http.StatusOK, transferReq)
}
