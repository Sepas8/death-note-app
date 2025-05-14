package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Sepas8/death-note-app/backend/api"
	"github.com/Sepas8/death-note-app/backend/models"
	"github.com/gorilla/mux"
)

func (s *Server) handleCreateKillDto(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)

	// Convertir ID a uint
	idStr := vars["id"]
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}
	id := uint(id64)

	// Verificar si ya existe un registro de muerte
	existingKill, err := s.KillRepository.FindByID(id)
	if existingKill != nil && existingKill.DeathExecuted {
		s.HandleError(w, http.StatusConflict, r.URL.Path, fmt.Errorf("person already dead"))
		return
	}

	var req api.KillRequestDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	person, err := s.PeopleRepository.FindByID(id)
	if err != nil || person == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path, fmt.Errorf("person not found"))
		return
	}

	// Calcular tiempo de muerte
	var deathTime time.Time
	duration := s.Config.DefaultKillDuration // 40 segundos

	if req.Cause != "" { // Cambiado de Description a Cause
		duration = s.Config.ExtendedKillDuration // 400 segundos
	}

	deathTime = time.Now().Add(duration)

	kill := &models.Kill{
		PersonID:     person.ID,
		Person:       *person, // Desreferenciar el puntero
		CauseOfDeath: req.Cause,
		TimeOfDeath:  deathTime,
	}

	// Programar la muerte
	s.taskQueue.StartTask(int(id), duration, func(k *models.Kill) error {
		k.DeathExecuted = true
		_, err := s.KillRepository.Save(k)
		return err
	}, kill)

	response := kill.ToKillTaskResponseDto()
	s.respondWithJSON(w, http.StatusCreated, response)
	s.logger.Info(http.StatusCreated, r.URL.Path, start)
}

func (s *Server) respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
