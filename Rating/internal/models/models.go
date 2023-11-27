package models

import (
	"github.com/gofrs/uuid"
)

type Rating struct {
	Id      *uuid.UUID `json:"id"`
	User_id *uuid.UUID `json:"user_id"`
	Song_id *uuid.UUID `json:"song_id"`
	Content string     `json:"content"`
}
