package models

import (
	"time"
)

type Thread struct {
	IdThread     int       `json:"id_thread"`
	Name         string    `json:"name"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	MessageCount int       `json:"message_count"`
	LikeCount    int       `json:"like_count"`
	CategoryName string    `json:"category_name"`
	FkIdCategory int       `json:"fk_id_category"`
	Author       *string   `json:"author"`
	FkIdUser     int       `json:"fk_id_user"`
	Status       int       `json:"status"`
}
