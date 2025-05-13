package api

type PersonRequestDto struct {
	Nombre string `json:"name"`
	Edad   int32  `json:"age"`
}

type PersonResponseDto struct {
	ID            int    `json:"person_id"`
	Nombre        string `json:"name"`
	Edad          int    `json:"age"`
	FechaCreacion string `json:"created_at"`
	Estado        string `json:"status"`
}

type ErrorResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Message     string `json:"message"`
}
