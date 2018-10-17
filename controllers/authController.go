package controllers

import (
	"encoding/json"
	u "AddressBook/utils"
	"net/http"
	"AddressBook/models"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Falha na solicitação"))
		return
	}

	resp := account.Create()
	u.Respond(w, resp)
}