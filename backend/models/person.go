package models

import (
	"backend-avanzada/api"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func (p *Person) ToPersonResponseDto() *api.PersonResponseDto {
	return &api.PersonResponseDto{
		ID:            int(p.ID),
		Nombre:        p.Name,
		Edad:          p.Age,
		FechaCreacion: p.CreatedAt.String(),
	}
}
