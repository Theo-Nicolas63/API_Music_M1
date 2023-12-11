package models

import (
	"github.com/gofrs/uuid"
)

type Song struct {
	Id     *uuid.UUID `json:"id"`
	Name   string     `json:"name"`
	Singer string     `json:"singer"`
}
