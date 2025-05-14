package repository

import (
	"github.com/Sepas8/death-note-app/backend/models"
	"gorm.io/gorm"
	"time"
)

type KillRepository struct {
	db *gorm.DB
}

func NewKillRepository(db *gorm.DB) *KillRepository {
	return &KillRepository{db: db}
}

func (k *KillRepository) ScheduleKill(personID uint, cause string, details string, executeAt time.Time) (*models.Kill, error) {
	kill := &models.Kill{
		PersonID:     personID,
		CauseOfDeath: cause,
		DeathDetails: details,
		TimeOfDeath:  executeAt,
	}

	err := k.db.Create(kill).Error
	if err != nil {
		return nil, err
	}

	return kill, nil
}

func (k *KillRepository) GetPendingKills() ([]*models.Kill, error) {
	var kills []*models.Kill
	now := time.Now()

	err := k.db.Where("time_of_death <= ? AND death_executed = ?", now, false).
		Preload("Person").
		Find(&kills).Error

	return kills, err
}

func (k *KillRepository) FindById(id int) (*models.Kill, error) {
	var kill models.Kill
	err := k.db.Where("id = ?", id).First(&kill).Error
	if err != nil {
		return nil, err
	}
	return &kill, nil
}

func (k *KillRepository) Save(kill *models.Kill) (*models.Kill, error) {
	err := k.db.Save(kill).Error
	if err != nil {
		return nil, err
	}
	return kill, nil
}

func (r *KillRepository) FindByID(id uint) (*models.Kill, error) {
	var kill models.Kill
	err := r.db.First(&kill, id).Error
	if err != nil {
		return nil, err
	}
	return &kill, nil
}
