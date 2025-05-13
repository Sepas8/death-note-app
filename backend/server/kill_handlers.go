func (s *Server) handleCreateKill(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	// Verificar si ya existe un registro de muerte
	existingKill, err := s.KillRepository.FindById(id)
	if existingKill != nil && existingKill.DeathExecuted {
		s.HandleError(w, http.StatusConflict, r.URL.Path, fmt.Errorf("person already dead"))
		return
	}

	var req api.KillRequestDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	person, err := s.PeopleRepository.FindById(id)
	if err != nil || person == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path, fmt.Errorf("person not found"))
		return
	}

	// Calcular tiempo de muerte
	var deathTime time.Time
	var duration time.Duration

	if req.Description == "" {
		// Muerte en 40 segundos
		duration = time.Duration(s.Config.KillDuration) * time.Second
	} else {
		// Muerte en 6 minutos 40 segundos (400 segundos)
		duration = time.Duration(s.Config.KillDurationWithDescription) * time.Second
	}

	deathTime = time.Now().Add(duration)

	kill := &models.Kill{
		PersonID:     person.ID,
		Person:       person,
		CauseOfDeath: req.Description,
		TimeOfDeath:  deathTime,
	}

	// Programar la muerte
	s.taskQueue.StartTask(id, duration, func(k *models.Kill) error {
		k.DeathExecuted = true
		_, err := s.KillRepository.Save(k)
		return err
	}, kill)

	response := kill.ToKillTaskResponseDto()
	respondWithJSON(w, http.StatusCreated, response)
	s.logger.Info(http.StatusCreated, r.URL.Path, start)
}