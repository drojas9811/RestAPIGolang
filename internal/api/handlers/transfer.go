package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/anthdm/gobank/internal/models"
)

func Transfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := new(model.TransferRequest) //add model
	if err := json.NewDecoder(r.Body).Decode(transferReq); err != nil {
		return err
	}
	defer r.Body.Close()

	return WriteJSON(w, http.StatusOK, transferReq)
}
