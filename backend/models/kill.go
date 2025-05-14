package models

import (
	"time"

	"github.com/Sepas8/death-note-app/backend/api"
	"gorm.io/gorm"
)

type Kill struct {
	gorm.Model
	PersonID      uint   `gorm:"not null"`
	Person        Person `gorm:"foreignKey:PersonID"`
	CauseOfDeath  string `gorm:"default:'heart attack'"`
	DeathDetails  string
	TimeOfDeath   time.Time
	DeathExecuted bool `gorm:"default:false"`
}

func (k *Kill) ToKillResponseDto() *api.KillResponseDto {
	return &api.KillResponseDto{
		Person:      k.Person.ToPersonResponseDto(),
		Description: k.CauseOfDeath,
		HoraMuerte:  k.TimeOfDeath.Format(time.RFC3339),
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
