package dto

type DtoMessageRequest struct {
	Content    string `json:"content"`
	FkIdUser   int    `json:"fk_id_user"`
	FkIdThread int    `json:"fk_id_thread"`
}

type DtoMessageResponse struct {
	IdMessage int     `json:"id_message"`
	Content   string  `json:"content"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	LikeCount int     `json:"likes"`
	Author    *string `json:"author"`
	FkIdUser  int     `json:"fk_id_user"`
}
