package dto

type DtoThreadRequest struct {
	Name         string `json:"name"`
	Content      string `json:"content"`
	Author       string `json:"author"`
	FkIdUser     int    `json:"fk_id_user"`
	FkIdCategory int    `json:"fk_id_category"`
}

type DtoThreadResponse struct {
	IdThread  int    `json:"id_thread"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	LikeCount int    `json:"likes"`
	Author    string `json:"author"`
}
