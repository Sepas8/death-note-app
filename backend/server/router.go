package server // Esta línea es esencial y debe ser la primera

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router configura y devuelve el enrutador HTTP
func (s *Server) Router() http.Handler {
	router := mux.NewRouter()
	router.Use(s.logger.RequestLogger)

	// Personas
	router.HandleFunc("/people", s.handlePeople).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/people/{id}", s.handlePeopleWithId).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	// Muertes
	router.HandleFunc("/kills", s.handleKills).Methods(http.MethodGet)
	router.HandleFunc("/people/{id}/death", s.handleCreateKillDto).Methods(http.MethodPost)

	return router
}
func (s *Server) handlePeopleWithId(w http.ResponseWriter, r *http.Request) {
	s.respondWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Not implemented"})
}

func (s *Server) handleKills(w http.ResponseWriter, r *http.Request) {
	s.respondWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Not implemented"})
}
