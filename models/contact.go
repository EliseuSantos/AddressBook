package models

import (
	u "AddressBook/utils"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	UserId uint `json:"user_id"`
}

func (contact *Contact) Validate() (map[string] interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "Nome do contato não pode ser vazio"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Número do contato não pode ser vazio"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "Usuário não encontrado"), false
	}

	return u.Message(true, "success"), true
}

func (contact *Contact) Create() (map[string] interface{}) {

	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}