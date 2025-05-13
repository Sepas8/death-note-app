package models

import (
	"backend/api"
	"time"

	"gorm.io/gorm"
)

type kill struct {
	gorm.Model
	PersonID      uint    `gorm:"not null"`
	Person        *Person `gorm:"foreignKey:PersonID"`
	CauseOfDeath  string  `gorm:"default:'heart attack'"`
	DeathDetails  string
	TimeOfDeath   time.Time
	DeathExecuted bool `gorm:"default:false"`
}

func (k *Kill) ToKillResponseDto() *api.KillResponseDto {
	return &api.KillResponseDto{
		Person:      k.Person.ToPersonResponseDto(),
		Description: k.CauseOfDeath,
	}
}

func (k *Kill) ToKillTaskResponseDto() *api.KillTaskResponseDto {
	status := "Pending"
	if k.DeathExecuted {
		status = "Executed"
	}
	return &api.KillTaskResponseDto{
		Person: k.Person.ToPersonResponseDto(),
		Status: status,
	}
}
