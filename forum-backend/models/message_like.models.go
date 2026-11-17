package models

type MessageLike struct {
	IdMessageLike int `json:"id_message_like"`
	FkIdMessage   int `json:"fk_id_message"`
	FkIdUser      int `json:"fk_id_user"`
}
