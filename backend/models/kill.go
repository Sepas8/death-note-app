package models

import (
	"backend-avanzada/api"

	"gorm.io/gorm"
)

type Kill struct {
	gorm.Model
	Description string
	PersonId    uint
	Person      *Person
}

func (k *Kill) ToKillResponseDto() *api.KillResponseDto {
	return &api.KillResponseDto{
		Person:      k.Person.ToPersonResponseDto(),
		Description: k.Description,
	}
}
