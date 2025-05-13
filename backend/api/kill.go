package api

type KillRequestDto struct {
	Description string `json:"description"`
}

type KillResponseDto struct {
	Person      *PersonResponseDto `json:"person"`
	Description string             `json:"description"`
}

type KillTaskResponseDto struct {
	Person *PersonResponseDto `json:"person"`
	Status string             `json:"status"`
}
