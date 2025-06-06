package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sepas8/death-note-app/backend/api"
	"github.com/Sepas8/death-note-app/backend/models"
)

func (s *Server) handlePeople(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetPeople(w, r)
	case http.MethodPost:
		s.handleCreatePerson(w, r)
	default:
		s.HandleError(w, http.StatusMethodNotAllowed, r.URL.Path, fmt.Errorf("method not allowed"))
	}
}

func (s *Server) handleGetPeople(w http.ResponseWriter, r *http.Request) {
	people, err := s.PeopleRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := make([]*api.PersonResponseDto, len(people))
	for i, p := range people {
		response[i] = p.ToPersonResponseDto()
	}

	s.respondWithJSON(w, http.StatusOK, response)
}

func (s *Server) handleCreatePerson(w http.ResponseWriter, r *http.Request) {
    var req api.PersonRequestDto
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
        return
    }

    // Debugging: Verifica los datos recibidos
    fmt.Printf("Received request: %+v\n", req)

    // Mapea los datos del DTO al modelo
    person := &models.Person{
        Name:     req.Nombre,
        Age:      req.Edad,      // Asigna el campo Edad
        PhotoURL: req.FotoURL,
    }

    // Crea la persona en el repositorio
    if err := s.PeopleRepository.Create(person); err != nil {
        s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
        return
    }

    // Responde con el DTO de respuesta
    s.respondWithJSON(w, http.StatusCreated, person.ToPersonResponseDto())
}