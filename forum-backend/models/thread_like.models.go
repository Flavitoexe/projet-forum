package models

type ThreadLike struct {
	IdThreadLike int `json:"id_thread_like"`
	FkIdThread   int `json:"fk_id_thread"`
	FkIdUser     int `json:"fk_id_user"`
}
