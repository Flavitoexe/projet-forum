package dto

type LoginResponseDto struct {
	IdUser      int    `json:"id_user"`
	Type        string `json:"type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
