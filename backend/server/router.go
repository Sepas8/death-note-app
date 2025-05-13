func (s *Server) router() http.Handler {
	router := mux.NewRouter()
	router.Use(s.logger.RequestLogger)
	
	// Personas
	router.HandleFunc("/people", s.HandlePeople).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/people/{id}", s.HandlePeopleWithId).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	
	// Muertes
	router.HandleFunc("/kills", s.HandleKills).Methods(http.MethodGet)
	router.HandleFunc("/people/{id}/death", s.HandleKillsWithId).Methods(http.MethodPost) // Cambiado a ruta más semántica
	
	return router
}