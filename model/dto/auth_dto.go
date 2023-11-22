package dto

type AuthRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Token string `json:"token"`
}
