package api

type PersonRequestDto struct {
	Nombre  string `json:"name" validate:"required"`
	Edad    int32  `json:"age"`
	FotoURL string `json:"photo_url" validate:"required,url"`
}

type PersonResponseDto struct {
	ID            int    `json:"person_id"`
	Nombre        string `json:"name"`
	Edad          int32  `json:"age"`
	FotoURL       string `json:"photo_url"`
	FechaCreacion string `json:"created_at"`
	Estado        string `json:"status"` // alive, pending, dead
}

type ErrorResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Message     string `json:"message"`
}
