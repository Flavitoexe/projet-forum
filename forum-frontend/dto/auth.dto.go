package dto

type RegisterRequestDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	IdUser int `json:"id_user"`
}
