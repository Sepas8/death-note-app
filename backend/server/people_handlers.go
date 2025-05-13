package server

import (
	"backend-avanzada/api"
	"backend-avanzada/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) HandlePeople(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllPeople(w, r)
		return
	case http.MethodPost:
		s.handleCreatePerson(w, r)
		return
	}
}

func (s *Server) HandlePeopleWithId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetPersonById(w, r)
		return
	case http.MethodPut:
		s.handleEditPerson(w, r)
		return
	case http.MethodDelete:
		s.handleDeletePerson(w, r)
		return
	}
}

func (s *Server) handleGetAllPeople(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	result := []*api.PersonResponseDto{}
	people, err := s.PeopleRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	for _, v := range people {
		alive := false
		kill, err := s.KillRepository.FindById(int(v.ID))
		if kill == nil && err == nil {
			alive = true
		}
		dto := &api.PersonResponseDto{
			ID:            int(v.ID),
			Nombre:        v.Name,
			Edad:          v.Age,
			FechaCreacion: v.CreatedAt.String(),
		}
		if alive {
			dto.Estado = "Vivo"
		} else {
			dto.Estado = "Muerto"
		}
		result = append(result, dto)
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

func (s *Server) handleGetPersonById(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}
	p, err := s.PeopleRepository.FindById(int(id))
	if p == nil && err == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path, fmt.Errorf("person with id %d not found", id))
		return
	}
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	resp := &api.PersonResponseDto{
		ID:            int(p.ID),
		Nombre:        p.Name,
		Edad:          p.Age,
		FechaCreacion: p.CreatedAt.String(),
	}
	response, err := json.Marshal(resp)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	s.logger.Info(http.StatusOK, r.URL.Path, start)
}

func (s *Server) handleCreatePerson(w http.ResponseWriter, r *http.Request) {
	var p api.PersonRequestDto
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	person := &models.Person{
		Name: p.Nombre,
		Age:  int(p.Edad),
	}
	person, err = s.PeopleRepository.Save(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pResponse := &api.PersonResponseDto{
		ID:            int(person.ID),
		Nombre:        person.Name,
		Edad:          person.Age,
		FechaCreacion: person.CreatedAt.String(),
	}
	result, err := json.Marshal(pResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func (s *Server) handleEditPerson(w http.ResponseWriter, r *http.Request) {
	var p api.PersonRequestDto
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
	person.Name = p.Nombre
	person.Age = int(p.Edad)
	person, err = s.PeopleRepository.Save(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pResponse := &api.PersonResponseDto{
		ID:            int(person.ID),
		Nombre:        person.Name,
		Edad:          person.Age,
		FechaCreacion: person.CreatedAt.String(),
	}
	result, err := json.Marshal(pResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(result)
}

func (s *Server) handleDeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
	err = s.PeopleRepository.Delete(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
