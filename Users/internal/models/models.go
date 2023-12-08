package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Id         *uuid.UUID `json:"Id"`         // identifiant de l'utilisateur
	MusicLiked int        `json:"MusicLiked"` // nombre de musiques aimé
	Name       string     `json:"Name"`       // nom de l'utilisateur
	Playlists  []string   `json:"Playlists"`  // tableau contenant le nom de toute les playlist de l'utilisateur
	Following  []string   `json:"Following"`  // tableau contenant les differentes artistes suivis par l'utilisateur
	Content    string     `json:"content"`
}
