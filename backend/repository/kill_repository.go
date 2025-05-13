// Métodos adicionales para KillRepository
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