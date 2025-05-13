package models

import (
	"import github.com/Sepas8/death-note-app/backend/api"
	"time"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name      string    `gorm:"not null"`       // Cambiar a FullName
	Age       int32     `gorm:"default:0"`      // Mantener
	PhotoURL  string    `gorm:"not null"`       // Nuevo campo requerido
	CreatedAt time.Time `gorm:"autoCreateTime"` // Mantener
}

// Actualizar el método ToPersonResponseDto
func (p *Person) ToPersonResponseDto() *api.PersonResponseDto {
	status := "Alive"
	// Verificar si existe registro en Kill para determinar estado
	return &api.PersonResponseDto{
		ID:            int(p.ID),
		Nombre:        p.Name,
		Edad:          p.Age,
		FotoURL:       p.PhotoURL, // Nuevo campo
		FechaCreacion: p.CreatedAt.Format(time.RFC3339),
		Estado:        status,
	}
}