func (s *Server) handleCreatePerson(w http.ResponseWriter, r *http.Request) {
	var p api.PersonRequestDto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	// Validar que la foto esté presente
	if p.FotoURL == "" {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, fmt.Errorf("photo is required"))
		return
	}

	person := &models.Person{
		Name:     p.Nombre,
		Age:      p.Edad,
		PhotoURL: p.FotoURL,
	}

	savedPerson, err := s.PeopleRepository.Save(person)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	// Programar muerte automática en 40 segundos
	kill := &models.Kill{
		PersonID:     savedPerson.ID,
		CauseOfDeath: "heart attack", // Valor por defecto
		TimeOfDeath:  time.Now().Add(time.Duration(s.Config.KillDuration) * time.Second),
	}

	// Usar el taskQueue para programar la muerte
	s.taskQueue.StartTask(int(savedPerson.ID), 
		time.Duration(s.Config.KillDuration)*time.Second,
		func(k *models.Kill) error {
			k.DeathExecuted = true
			_, err := s.KillRepository.Save(k)
			return err
		}, 
		kill)

	respondWithJSON(w, http.StatusCreated, savedPerson.ToPersonResponseDto())
}