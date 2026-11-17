package models

import (
	"time"
)

type Message struct {
	IdMessage  int        `json:"id_message"`
	Content    string     `json:"content"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	LikeCount  int        `json:"likes"`
	FkIdUser   *int       `json:"fk_id_user"`
	Author     *string    `json:"author"`
	FkIdThread int        `json:"fk_id_thread"`
}
