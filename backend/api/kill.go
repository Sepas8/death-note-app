package api

type KillRequestDto struct {
	Cause    string `json:"cause"`             // Requerido dentro de 40 segundos
	Detalles string `json:"details,omitempty"` // Opcional dentro de 6:40 min adicionales
}

type KillResponseDto struct {
	Person      *PersonResponseDto `json:"person"`
	Description string             `json:"description"`
	HoraMuerte  string             `json:"death_time"`
}

type KillTaskResponseDto struct {
	Person *PersonResponseDto `json:"person"`
	Status string             `json:"status"` // pending, executing, completed
}
