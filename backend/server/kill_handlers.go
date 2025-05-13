package server

import (
	"backend/api"
	"backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) HandleKills(w http.ResponseWriter, r *http.Request) {
	s.handleGetAllKills(w, r)
}

func (s *Server) HandleKillsWithId(w http.ResponseWriter, r *http.Request) {
	s.handleCreateKill(w, r)
}

func (s *Server) handleGetAllKills(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	result := []*api.KillResponseDto{}
	kills, err := s.KillRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	for _, v := range kills {
		result = append(result, v.ToKillResponseDto())
	}
	response, err := json.Marshal(result)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	s.logger.Info(http.StatusOK, r.URL.Path, start)
}

func (s *Server) handleCreateKill(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var k api.KillRequestDto
	var duration time.Duration
	err := json.NewDecoder(r.Body).Decode(&k)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, exists := s.taskQueue.tasks[int(id)]
	if exists {
		s.HandleError(w, http.StatusConflict, r.URL.Path, fmt.Errorf("task with id %d is already in progress", id))
		return
	}
	person, err := s.PeopleRepository.FindById(int(id))
	if person == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	kill, err := s.KillRepository.FindById(int(id))
	if kill != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	kill = &models.Kill{
		Description: k.Description,
		PersonId:    person.ID,
		Person:      person,
	}
	if strings.Compare(k.Description, "") == 0 {
		duration = time.Duration(s.Config.KillDuration) * time.Second
	} else {
		duration = time.Duration(s.Config.KillDurationWithDescription) * time.Second
	}
	killFunc := func(k *models.Kill) error {
		_, err := s.KillRepository.Save(k)
		if err != nil {
			return err
		}
		return nil
	}
	s.taskQueue.StartTask(int(person.ID), duration, killFunc, kill)
	result, err := json.Marshal(&api.KillTaskResponseDto{
		Person: person.ToPersonResponseDto(),
		Status: "In progress.",
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
	s.logger.Info(http.StatusCreated, r.URL.Path, start)
}
