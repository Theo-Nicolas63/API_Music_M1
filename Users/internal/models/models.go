package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Id         *uuid.UUID `json:"Id"`         // identifiant de l'utilisateur
	Name       string     `json:"Name"`       // nom de l'utilisateur
	Username   string     `json:"Username"` 
	Password   string     `json: "Password"`
}
