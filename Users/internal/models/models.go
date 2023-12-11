package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Id         *uuid.UUID `json:"Id"`         // identifiant de l'utilisateur
	Name       string     `json:"Name"`       // nom de l'utilisateur
	MusicLiked int        `json:"MusicLiked"` // nombre de musiques aimés
}
